package utils

import (
	"net/url"
	"strings"
)

// CleanURL removes search parameters from an URL
func CleanURL(inputURL string) (string, error) {
	// if we can
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		return "", err
	}
	parsedURL.RawQuery = ""
	return parsedURL.String(), nil
}

// Extracts host name from a given URL. Keep the host name after `www.` for host names starting with it
func URLHostName(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	hostname := parsedURL.Hostname()
	if strings.HasPrefix(hostname, "wwww.") {
		hostname = hostname[4:]
	}
	return hostname, nil
}
