package pkgutils

import "testing"

func TestInArray_String(t *testing.T) {
	haystack := []string{"apple", "banana", "cherry"}

	tests := []struct {
		needle   string
		expected bool
	}{
		{"apple", true},
		{"banana", true},
		{"cherry", true},
		{"grape", false},
		{"", false},
	}

	for _, test := range tests {
		result := InArray(test.needle, haystack)
		if result != test.expected {
			t.Errorf("InArray(%q) = %v; want %v", test.needle, result, test.expected)
		}
	}
}

func TestInArray_Int(t *testing.T) {
	haystack := []int{1, 2, 3, 4, 5}

	tests := []struct {
		needle   int
		expected bool
	}{
		{1, true},
		{3, true},
		{5, true},
		{0, false},
		{6, false},
	}

	for _, test := range tests {
		result := InArray(test.needle, haystack)
		if result != test.expected {
			t.Errorf("InArray(%d) = %v; want %v", test.needle, result, test.expected)
		}
	}
}
