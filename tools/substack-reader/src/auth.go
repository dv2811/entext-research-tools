package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"entext-applications/internal/substack"
)

func runAuthHelp() {
	fmt.Println(`Authenticate with Substack via email link.

Usage: substack auth

Flow:
1. Enter your email address
2. Check your email for the authentication link
3. Copy and paste the link into the terminal
4. Session will be saved automatically`)
}

func runAuth(client *substack.Client, session *substack.Session, args []string) {
	reader := bufio.NewReader(os.Stdin)

	// Prompt for email
	fmt.Print("Enter your Substack email: ")
	emailAddr, _ := reader.ReadString('\n')
	emailAddr = strings.TrimSpace(emailAddr)

	if emailAddr == "" {
		fmt.Fprintf(os.Stderr, "Error: email is required\n")
		os.Exit(1)
	}

	// Validate email format
	if !strings.Contains(emailAddr, "@") {
		fmt.Fprintf(os.Stderr, "Error: invalid email format\n")
		os.Exit(1)
	}

	// Request authentication link from Substack
	fmt.Printf("\nRequesting authentication link for %s...\n", emailAddr)
	_, err := client.StartEmailLinkLogin(emailAddr, "https://substack.com")
	if err != nil {
		exitWithError(err)
	}

	fmt.Println("✓ Authentication link sent to your email!")
	fmt.Println()
	fmt.Println("Check your email inbox (and spam folder) for the link from Substack.")
	fmt.Println()

	// Prompt for authentication URL
	fmt.Print("Paste the authentication URL: ")
	authURL, _ := reader.ReadString('\n')
	authURL = strings.TrimSpace(authURL)

	if authURL == "" {
		fmt.Fprintf(os.Stderr, "Error: authentication URL is required\n")
		os.Exit(1)
	}

	if !strings.HasPrefix(authURL, "http") {
		fmt.Fprintf(os.Stderr, "Error: invalid URL format\n")
		os.Exit(1)
	}

	// Complete authentication
	fmt.Println("\nCompleting authentication...")
	err = client.AuthenticateFromResponse(session, authURL)
	if err != nil {
		exitWithError(err)
	}

	// Set email on session
	session.Email = emailAddr

	fmt.Println("✓ Authentication successful!")
	fmt.Println()
	fmt.Printf("Session saved for: %s\n", emailAddr)
}
