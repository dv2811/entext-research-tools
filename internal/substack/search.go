package substack

import (
	"errors"
	"fmt"
	"io"
	"net/url"
	"strconv"

	"entext-applications/internal/validator"
	"github.com/goccy/go-json"
)

var (
	ErrInvalidSearchMode = errors.New("substack: invalid search mode")
)

type SbkSearchFeedMode string

const (
	SearchModeAll SbkSearchFeedMode = "all"
	SearchModeTop                   = "top"
	SearchModeSub                   = "subscribed"
)

type SearchRequest struct {
	Query  string            `json:"query" jsonschema:"search query, required"`
	Mode   SbkSearchFeedMode `json:"mode" jsonschema:"search mode for Substack content, One of: 'top', 'all','subscribed'"`
	Lang   string            `json:"language" jsonschema:"language for substack search, in 2-letter country code format, default en,optional"`
	Cursor string            `json:"cursor,omitempty" jsonschema:"pagination cursor for retrieving the next set off top post results"`
	Page   int               `json:"page" jsonschema:"zero-indexed pagination number, maximum 10. not used for top post search"`
}

func (req *SearchRequest) Validate(v *validator.Validator) {
	v.Check(req.Query != "", "query", "cannot be empty")
	v.Check(validator.In(req.Mode, "top", "all", "subscribed"), "mode", "invalid search mode value")

	// handle malformatted cases
	switch req.Mode {
	case "top":
		v.Check(req.Page == 0, "page", "page must not be set for top result search")
	case "all", "subscribed":
		v.Check(req.Lang == "" || len(req.Lang) == 2, "language", "invalid language code")
		v.Check(req.Cursor == "", "cursor", "cursor not allowed for this search mode")
		v.Check(req.Page <= 10 && req.Page > -1, "page", "invalid page number")
	}
}

type SearchResponse struct {
	Results []Post `json:"results" jsonschema:"search results"`
	HasMore bool   `json:"more" jsonschema:"boolean tag informing whether there are more results"`
	Cursor  string `json:"cursor,omitempty" jsonschema:"pagination cursor for retrieving the next set of top post results"`
}

// SearchPosts provides a single entry point for different substack content seach modes
func (c *Client) SearchPosts(session *Session, req SearchRequest) (SearchResponse, error) {
	v := url.Values{}
	v.Set("query", req.Query)

	// adjust URL query based on mode
	var pathMode string
	switch req.Mode {

	case "top":
		pathMode = "top"
		v.Set("fromSuggestedSearch", "false")

	case "subscribed":
		pathMode = "post"
		v.Set("filter", "subscribed")
		v.Set("page", strconv.Itoa(req.Page))
		if req.Lang == "" {
			v.Set("language", "en")
		} else {
			v.Set("language", req.Lang)
		}
	case "all":
		pathMode = "post"
		v.Set("includePlatformResults", "true")
		v.Set("filter", "all")
		v.Set("page", strconv.Itoa(req.Page))
		if req.Lang == "" {
			v.Set("language", "en")
		} else {
			v.Set("language", req.Lang)
		}

	default:
		return SearchResponse{}, ErrInvalidSearchMode
	}

	queryURL := fmt.Sprintf("https://substack.com/api/v1/%s/search?%s", pathMode, v.Encode())

	headers := map[string][]string{
		"Accept":          {"*/*"},
		"Accept-Language": {"en-US,en;q=0.9"},
		"user-agent":      {"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36"},
		"Referer":         {"https://substack.com/"},
		"Sec-Fetch-Dest":  {"empty"},
		"Sec-Fetch-Mode":  {"cors"},
		"Sec-Fetch-Site":  {"same-origin"},
	}

	err := session.AuthorizeHeaders(headers)
	if err != nil {
		return SearchResponse{}, err
	}

	resp, err := c.getResponse("GET", queryURL, nil, headers)
	if err != nil {
		return SearchResponse{}, err
	}

	defer func() {
		io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
	}()

	// optional response update
	_ = session.Refresh(resp)

	searchResp := SearchResponse{}

	if req.Mode != "top" {
		err = json.NewDecoder(resp.Body).Decode(&searchResp)
		if err != nil {
			return SearchResponse{}, err
		}

		// sort posts for better organization
		SortPosts(searchResp.Results)

	} else {
		err = parseTopPostContent(resp.Body, &searchResp)
		if err != nil {
			return SearchResponse{}, err
		}
	}
	return searchResp, nil
}

type TopPostResult struct {
	Items  []SearchItem `json:"items"`
	Cursor string       `json:"nextCursor"`
}

type SearchItem struct {
	Type string `json:"type"`
	Post *Post  `json:"post"`
}

// parseTopPostContent decodes top post content and transforms it to target format
func parseTopPostContent(r io.Reader, content *SearchResponse) error {
	result := TopPostResult{}
	err := json.NewDecoder(r).Decode(&result)
	if err != nil {
		return err
	}

	posts := make([]Post, 0, len(result.Items))

	for _, item := range result.Items {
		if item.Post == nil && item.Type != "post" {
			continue
		}
		posts = append(posts, *item.Post)
	}

	// set content fields
	content.Results = posts
	content.Cursor = result.Cursor
	content.HasMore = result.Cursor != ""
	return nil
}
