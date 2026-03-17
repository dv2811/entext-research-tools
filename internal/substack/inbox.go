package substack

import (
	"bytes"
	"entext-applications/internal/utils"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/url"
	"slices"

	"github.com/goccy/go-json"
)

type Post struct {
	Title      string      `json:"title" jsonschema:"Title of the post"`
	Subtitle   string      `json:"subtitle" jsonschema:"Post's subtitle providing extra context"`
	PostDate   string      `json:"post_date" jsonschema:"Publication date of the post,pattern=^\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}\\.\\d{3}Z$"`
	Url        string      `json:"canonical_url" jsonschema:"Canonical URL of the post, can be used for viewing article content if article content endpoint fails,maxLength=1000"`
	PubId      uint32      `json:"publication_id" jsonschema:"ID of the publication,minimum=1"`
	ID         uint32      `json:"id" jsonschema:"Unique identifier of the post,minimum=1"`
	BodyHTML   string      `json:"body_html,omitempty" jsonschema:"JSON escaped HTML body of the post,maxLength=50000"` // HTML body of a post
	AudioItems []AudioItem `json:"audio_items,omitempty" jsonschema:"Audio items associated with the post"`
	ByLines    []ByLine    `json:"publishedBylines,omitempty" jsonschema:"Authors of the post"`
}

// we don't have use for ByLine other than showing author name
type ByLine struct {
	Name string `json:"name,omitempty" jsonschema:"Name of the author"`
}

type AudioItem struct {
	AudioURl string `json:"audio_url" jsonschema:"Audio URL for podcast or text-to-speech,maxLength=1000"` // audio URL for podcast of text-to-speech audio
}

type Inbox struct {
	Posts          []Post `json:"posts" jsonschema:"List of posts in the inbox"`
	NextTimeCursor string `json:"next_time_cursor" jsonschema:"Timestamp of the earliest document in the current batch, used for retrieving the next batch of posts,pattern=^\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}\\.\\d{3}Z$"` // timestamp of the earliest document in the current batch
}

// SortPosts perform in place sorting for each post result, in reverse chronological order
func SortPosts(posts []Post) {
	// var lastTimeCursor string
	if len(posts) > 0 {
		// sort by reverse datetime order, then choose last element for next time cursor
		slices.SortFunc(posts, func(a, b Post) int {
			if len(a.PostDate) > len(a.PostDate) {
				return -1
			} else if len(a.PostDate) < len(a.PostDate) {
				return 1
			} else if a.PostDate > b.PostDate {
				return -1
			} else if a.PostDate == b.PostDate {
				return 0
			}
			return 1
		})
	}
}

// GetChronologicalInbox gets chronologically sorted inbox item after a given timestamp
func (c *Client) GetChronologicalInbox(session *Session, lastTimeCursor string) (*Inbox, error) {
	v := url.Values{}
	v.Set("inboxType", "inbox")
	v.Set("limit", "20") // not modifiable
	if lastTimeCursor != "" {
		v.Set("after", lastTimeCursor)
	}

	queryURL := fmt.Sprintf("https://substack.com/api/v1/reader/posts?%s", v.Encode())
	headers := map[string][]string{
		"Accept-Encoding": {"application/json"},
		"user-agent":      {"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36"},
		"Origin":          {"https://substack.com"},
	}

	err := session.AuthorizeHeaders(headers)
	if err != nil {
		return nil, err
	}

	// get response
	resp, err := c.getResponse("GET", queryURL, nil, headers)
	if err != nil {
		return nil, err
	}

	// clean up resources
	defer func() {
		io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
	}()

	// session refresh - ignore error
	_ = session.Refresh(resp)

	inboxItems := &Inbox{}
	err = json.NewDecoder(resp.Body).Decode(inboxItems)
	if err != nil {
		return nil, err
	}

	// sort posts
	SortPosts(inboxItems.Posts)

	// Set the NextTimeCursor based on the earliest post in the batch
	count := len(inboxItems.Posts)
	if count > 0 {
		inboxItems.NextTimeCursor = inboxItems.Posts[count-1].PostDate
	}

	return inboxItems, nil
}

// Substack article content
// https://substack.com/app-link/post?publication_id=47580&post_id=186860256
// https://www.substack.com/api/v1/posts/by-id/145856537

type Article struct {
	Post Post `json:"post" jsonschema:"The post content"`
}

func (c *Client) GetArticleContent(session *Session, postID uint32, baseURL string) (*Article, error) {
	var queryURL string
	if baseURL == "" {
		queryURL = fmt.Sprintf("https://www.substack.com/api/v1/posts/by-id/%d", postID)
	} else {
		queryURL = fmt.Sprintf("https://%s/api/v1/posts/by-id/%d", baseURL, postID)
	}
	headers := map[string][]string{
		"Accept-Encoding": {"application/json"},
		"user-agent":      {"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36"},
		"Origin":          {"https://substack.com"},
	}

	err := session.AuthorizeHeaders(headers)
	if err != nil {
		return nil, err
	}

	// get response
	resp, err := c.getResponse("GET", queryURL, nil, headers)
	if err != nil {
		return nil, err
	}

	// clean up resources
	defer func() {
		io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
	}()

	// session refresh - ignore error
	_ = session.Refresh(resp)

	article := &Article{}
	err = json.NewDecoder(resp.Body).Decode(article)
	if err != nil {
		return nil, err
	}
	// article.Post.BodyHTML = reduceHTML(article.Post.BodyHTML)
	return article, nil
}

var isVoid = map[string]bool{
	"area":   true,
	"base":   true,
	"br":     true,
	"col":    true,
	"embed":  true,
	"hr":     true,
	"img":    true,
	"input":  true,
	"link":   true,
	"meta":   true,
	"param":  true,
	"source": true,
	"track":  true,
	"wbr":    true,
}

// reduceHTML shortens bodyHTML by simplifying image tags
func reduceHTML(bodyHTML string) string {
	convertedBytes := utils.StringToBytes(bodyHTML)
	buffer := convertedBytes[:0]

	tkn := html.NewTokenizer(bytes.NewReader(convertedBytes))
	excludeTagDepth := 0
	imageURL := make([]byte, 0, 256)

	for tokenType := tkn.Next(); tkn.Err() == nil; tokenType = tkn.Next() {
		if excludeTagDepth > 0 {
			if tokenType == html.StartTagToken {
				excludeTagDepth++
				tagName, hasAttr := tkn.TagName()
				if isVoid[string(tagName)] {
					excludeTagDepth--
				}
				if hasAttr && string(tagName) == "a" {
					for {
						key, val, more := tkn.TagAttr()
						if string(key) == "href" {
							imageURL = append(imageURL[:0], val...)
							break
						}
						if !more {
							break
						}
					}
				}
			} else if tokenType == html.EndTagToken {
				excludeTagDepth--
				if excludeTagDepth == 0 && len(imageURL) > 0 {
					buffer = append(buffer, `<img src="`...)
					buffer = append(buffer, imageURL...)
					buffer = append(buffer, `" alt="" />`...)
				}
			}
			continue
		}

		rawToken := tkn.Raw()
		if bytes.HasPrefix(rawToken, []byte(`<button`)) ||
			bytes.HasPrefix(rawToken, []byte(`<svg`)) ||
			(bytes.HasPrefix(rawToken, []byte(`<div`)) && bytes.Contains(rawToken, []byte(`captioned-image-container`))) {
			excludeTagDepth = 1
			continue
		}
		buffer = append(buffer, rawToken...)
	}
	return string(buffer)
}
