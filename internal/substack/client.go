package substack

import (
	"net/http"
	"net/http/cookiejar"

	"bytes"
	"errors"
	"io"
	"sync"

	"entext-applications/internal/utils"
	"github.com/goccy/go-json"
)

var bufferPool = sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}

// define basic errors
var (
	ErrLoginFailed = errors.New("login failed")
	ErrAPIError    = errors.New("error: api error")
)

// Client provides a wrapper around http.Client from which http requests will be run
type Client struct {
	client *http.Client
}

func NewClient() *Client {
	client := utils.NewClient()
	// cookie will be set manually in header, client will be shared across users
	jar, _ := cookiejar.New(nil)
	client.Jar = jar

	return &Client{
		client: client,
	}
}

func (c *Client) getResponse(method, url string, requestBody io.Reader, headers http.Header) (*http.Response, error) {
	// prepare request
	req, err := http.NewRequest(method, url, requestBody)
	if err != nil {
		return nil, err
	}

	if len(headers) > 0 {
		req.Header = headers
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	// HTTP error handling
	if resp.StatusCode > 399 {
		io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
		return nil, errors.Join(ErrAPIError, errors.New(resp.Status))
	}

	return resp, nil
}

type LoginParamters struct {
	Redirect string `json:"redirect,omitempty"`
	ForPub   string `json:"for_pub"`
	Email    string `json:"email"`
	// Password can be omitted during email login link flow
	Password string `json:"password,omitempty"`
	Captcha  any    `json:"captcha_response"`
}

// EmailLinkResult contains the authentication URL and callback info
type EmailLinkResult struct {
	Email      string
	AuthURL    string
	CallbackCh chan string
}

// StartEmailLinkLogin triggers a login by email link/code process
// Returns the authentication URL that will be sent to the user's email
func (c *Client) StartEmailLinkLogin(email, redirectURL string) (string, error) {
	headers := map[string][]string{
		"Host":            {"substack.com"},
		"User-Agent":      {"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36"},
		"Accept":          {"*/*"},
		"Accept-Language": {"en-US,en;q=0.5"},
		"Accept-Encoding": {"application/json"},
		"Content-Type":    {"application/json"},
		"Content-Length":  {"82"},
		"Origin":          {"https://substack.com"},
		"DNT":             {"1"},
		"Sec-GPC":         {"1"},
		"Connection":      {"keep-alive"},
		"Referer":         {"https://substack.com/sign-in?redirect=%2F"},
		"Sec-Fetch-Dest":  {"empty"},
		"Sec-Fetch-Mode":  {"cors"},
		"Sec-Fetch-Site":  {"same-origin"},
	}

	loginInput := LoginParamters{
		Redirect: redirectURL,
		Email:    email,
		Captcha:  nil,
	}

	buf, _ := bufferPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer bufferPool.Put(buf)

	err := json.NewEncoder(buf).Encode(loginInput)
	if err != nil {
		return "", err
	}

	// making request to the API
	resp, err := c.getResponse("POST", "https://substack.com/api/v1/email-login", buf, headers)
	if err != nil {
		return "", errors.Join(ErrLoginFailed, err)
	}

	// HTTP response not OK
	if resp.StatusCode > 399 {
		io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
		return "", errors.Join(ErrLoginFailed, errors.New(resp.Status))
	}

	// Read response to get the auth URL
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
		return "", err
	}
	io.Copy(io.Discard, resp.Body)
	resp.Body.Close()

	// Extract the authentication URL from response
	authURL, _ := result["redirect_url"].(string)
	if authURL == "" {
		authURL, _ = result["url"].(string)
	}

	return authURL, nil
}

// PasswordLogin performs login to Substack using email and password combination
// Due to Cloudflare blocking, using this has a lower chance of success than using email login link flow
func (c *Client) PasswordLogin(email, password string) (*Session, error) {
	headers := map[string][]string{
		"Host":            {"substack.com"},
		"User-Agent":      {"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36"},
		"Accept":          {"*/*"},
		"Accept-Language": {"en-US,en;q=0.5"},
		"Accept-Encoding": {"application/json"},
		"Content-Type":    {"application/json"},
		"Content-Length":  {"82"},
		"Origin":          {"https://substack.com"},
		"DNT":             {"1"},
		"Sec-GPC":         {"1"},
		"Connection":      {"keep-alive"},
		"Referer":         {"https://substack.com/sign-in?redirect=%2F"},
		"Sec-Fetch-Dest":  {"empty"},
		"Sec-Fetch-Mode":  {"cors"},
		"Sec-Fetch-Site":  {"same-origin"},
	}

	loginInput := LoginParamters{
		Email:    email,
		Password: password,
		Captcha:  "",
	}

	buf, _ := bufferPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer bufferPool.Put(buf)

	err := json.NewEncoder(buf).Encode(loginInput)
	if err != nil {
		return nil, err
	}

	// making request to the API
	resp, err := c.getResponse("POST", "https://substack.com/api/v1/login", buf, headers)
	if err != nil {
		return nil, errors.Join(ErrLoginFailed, err)
	}

	// HTTP response not OK
	if resp.StatusCode != 200 {
		return nil, errors.Join(ErrLoginFailed, errors.New(resp.Status))
	}

	// define new session
	session := new(Session)
	session.Email = email

	// use HTTP response to populate session fields
	err = session.LoadFromResponse(resp)
	if err != nil {
		return nil, errors.Join(ErrLoginFailed, err)
	}

	// response body from this request is not important to us
	io.Copy(io.Discard, resp.Body)
	resp.Body.Close()

	return session, nil
}

// AuthenticatedURLFlow generates Substack session from authenticated login link
func (c *Client) AuthenticateFromResponse(session *Session, url string) error {
	resp, err := c.getResponse("GET", url, nil, nil)
	if err != nil {
		return errors.Join(ErrLoginFailed, err)
	}

	// check HTTP errors
	if resp.StatusCode != 200 {
		return errors.Join(ErrLoginFailed, errors.New(resp.Status))
	}

	// use HTTP response to populate session fields
	err = session.LoadFromResponse(resp)
	if err != nil {
		return errors.Join(ErrLoginFailed, err)
	}

	// clean up resources to release connection back to the pool
	io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	return nil
}
