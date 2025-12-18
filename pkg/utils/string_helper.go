package pkgutils

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"strings"
	"time"
)

// Convert a time.Time to base64 string
func TimeToBase64(t time.Time) string {
	return base64.StdEncoding.EncodeToString([]byte(t.Format(time.RFC3339Nano)))
}

// Convert a base64 string to time.Time
func Base64ToTime(s string) (time.Time, error) {
	bytes, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return time.Time{}, err
	}
	return time.Parse(time.RFC3339Nano, string(bytes))
}

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
