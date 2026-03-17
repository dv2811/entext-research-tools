package utils

import (
	"strings"
)

// ContainsAny checks if a string contains any of the substrings (case-insensitive)
func ContainsAny(s string, substrings []string) bool {
	for _, substr := range substrings {
		if strings.Contains(strings.ToLower(s), strings.ToLower(substr)) {
			return true
		}
	}
	return false
}

// In-place filtering a given slice to only unique values
func ApplyUniqueFilter[T comparable](values []T) []T {
	// no work to be done here
	if len(values) < 2 {
		return values
	}

	// unique filter
	isDuplicate := make(map[T]bool, len(values))
	output := values[:0]
	for indx, v := range values {
		if isDuplicate[v] {
			continue
			// assign seen part of the slice if no duplicate is found yet
		} else if indx+1 == len(isDuplicate) {
			output = values[:indx+1]
		} else {
			output = append(output, v)
		}
		isDuplicate[v] = true
	}
	return output
}
