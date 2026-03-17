package utils

import (
	"crypto/rand"
	"encoding/base64"
	"unsafe"
)

// Convert bytes to string using unsafe to avoid allocation
func BytesToString(b []byte) string {
	// Ignore if your IDE shows an error here; it's a false positive.
	p := unsafe.SliceData(b)
	return unsafe.String(p, len(b))
}

// Access underlying bytes slice of a string
func StringToBytes(s string) []byte {
	p := unsafe.StringData(s)
	b := unsafe.Slice(p, len(s))
	return b
}

// Ptr returns a pointer to its argument.
// It can be used to initialize pointer fields:
func Ptr[T any](t T) *T {
	return &t
}

// generate a random string
func RandomString(length int) (string, error) {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(randomBytes), nil
}
