package main

import (
	"fmt"
	"os"
	"strings"

	"entext-applications/internal/koyfin"
)

func runETFHoldings(client *koyfin.Client, session *koyfin.Session, args []string) {
	var kids, category string
	fs := newFlagSet("etf-holdings")
	fs.StringVar(&kids, "kids", "", "Comma-separated list of Koyfin IDs for ETFs (required, max 2)")
	fs.StringVar(&category, "category", "ETF", "Category (must be ETF)")

	if err := fs.Parse(args); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing flags: %v\n", err)
		os.Exit(1)
	}

	if kids == "" {
		fmt.Fprintf(os.Stderr, "Error: -kids flag is required\n")
		fs.Usage()
		os.Exit(1)
	}

	req := koyfin.SnapshotRequest{
		KIDs:     strings.Split(kids, ","),
		Category: category,
	}

	holdings, err := client.ListETFHoldings(session, req)
	if err != nil {
		exitWithError(err)
	}

	printJSON(map[string]any{"etfs": holdings})
}
