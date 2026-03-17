package main

import (
	"fmt"
	"os"
	"strconv"

	"entext-applications/internal/substack"
)

func runArticleHelp() {
	fmt.Println(`Get article content by post ID.

Usage: substack article [flags]

Flags:
  -post-id uint
        Post ID to retrieve content for (required)
  -base-url string
        Custom Substack domain (optional, e.g., 'substack.com')

Examples:
  substack article -post-id 123456
  substack article -post-id 123456 -base-url "substack.com"`)
}

func runArticle(client *substack.Client, session *substack.Session, args []string) {
	fs := newFlagSet("article")
	postID := fs.Uint("post-id", 0, "Post ID to retrieve content for (required)")
	baseURL := fs.String("base-url", "", "Custom Substack domain (optional, e.g., 'substack.com')")

	if err := fs.Parse(args); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing flags: %v\n", err)
		os.Exit(1)
	}

	if *postID == 0 {
		fmt.Fprintf(os.Stderr, "Error: -post-id flag is required\n")
		fs.Usage()
		os.Exit(1)
	}

	article, err := client.GetArticleContent(session, uint32(*postID), *baseURL)
	if err != nil {
		exitWithError(err)
	}

	printJSON(map[string]any{"data": article})
}

// Helper function to convert string to uint32
func parsePostID(idStr string) (uint32, error) {
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(id), nil
}
