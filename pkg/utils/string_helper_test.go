package utils

import (
	"testing"
)

func TestStringHelper(t *testing.T) {
	t.Run("Cfirst", func(t *testing.T) {
		tests := []struct {
			input    string
			expected string
		}{
			{"hello", "Hello"},
			{"world", "World"},
			{"", ""},
		}

		for _, test := range tests {
			result := Cfirst(test.input)
			if result != test.expected {
				t.Errorf("Cfirst(%q) = %q; want %q", test.input, result, test.expected)
			}
		}
	})

	t.Run("Ucfirst", func(t *testing.T) {
		tests := []struct {
			input    string
			expected string
		}{
			{"Hello", "hello"},
			{"World", "world"},
			{"", ""},
		}

		for _, test := range tests {
			result := Ucfirst(test.input)
			if result != test.expected {
				t.Errorf("Ucfirst(%q) = %q; want %q", test.input, result, test.expected)
			}
		}
	})

	t.Run("Trim", func(t *testing.T) {
		tests := []struct {
			input    string
			ch       string
			expected string
		}{
			{"Hello World", " ", "Hello World"},
			{"Hello World", "-", "Hello World"},
			{"", " ", ""},
		}

		for _, test := range tests {
			result := Trim(test.input, test.ch)
			if result != test.expected {
				t.Errorf("Trim(%q) = %q; want %q", test.input, result, test.expected)
			}
		}
	})
}
