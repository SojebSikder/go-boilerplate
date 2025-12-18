package pkgutils

import "slices"

func InArray[T comparable](needle T, haystack []T) bool {
	return slices.Contains(haystack, needle)
}
