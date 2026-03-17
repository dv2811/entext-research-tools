package validator

import (
	"net/url"
	"regexp"
	"strings"
)

var (
	EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

type Validator struct {
	Errors map[string]string
}

func New() *Validator {
	return &Validator{Errors: make(map[string]string)}
}
func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}
func (v *Validator) AddError(key, message string) {
	if _, ok := v.Errors[key]; !ok {
		v.Errors[key] = message
	}
}
func (v *Validator) Check(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}

// In returns true if a specific value is in a list of comparable values.
func In[T comparable](value T, list ...T) bool {
	for i := range list {
		if value == list[i] {
			return true
		}
	}
	return false
}

// Matches returns true if a string value matches a specific regexp pattern.
func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}

// Unique returns true if all string values in a slice are unique.
// in the context of this package - check if genres are unique values
func Unique(values []string) bool {
	uniqueValues := make(map[string]bool)
	for _, value := range values {
		uniqueValues[value] = true
	}
	return len(values) == len(uniqueValues)
}

func ContentTypeAllowed(contentType string, extensions []string) bool {
	var allowed bool
	for _, extension := range extensions {
		if strings.HasSuffix(contentType, extension) {
			allowed = true
			break
		}
	}
	return allowed
}

// check if a string is valid URL
func ValidURL(text string) bool {
	_, err := url.ParseRequestURI(text)
	if err != nil {
		return false
	}

	u, err := url.Parse(text)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}

// IsSubset returns true if all elements in src are contained in checker
func IsSubset[T comparable](src []T, checker []T) bool {
	checkSet := make(map[T]bool)
	for _, item := range checker {
		checkSet[item] = true
	}

	for _, item := range src {
		if !checkSet[item] {
			return false
		}
	}
	return true
}
