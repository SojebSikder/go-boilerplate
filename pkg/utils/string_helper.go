package utils

import (
	"crypto/rand"
	"encoding/hex"
	"strings"
)

// Generate a random string of a given length
func RandomString(length int) string {
	bytes := make([]byte, length)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

// Capitalize the first letter of a string
func Cfirst(s string) string {
	if len(s) == 0 {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// Uncapitalize the first letter of a string
func Ucfirst(s string) string {
	if len(s) == 0 {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

// Trim a string of a given character
func Trim(s string, ch string) string {
	if len(s) == 0 {
		return ""
	}
	return strings.Trim(s, ch)
}
