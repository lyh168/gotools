package sliceutils

import "strings"

// StringSliceIsEmpty Checks if an slice is nil or length equals 0
func StringSliceIsEmpty(slice []string) bool {
	return len(slice) <= 0
}

// StringSliceContains Checks if the value is in the given slice.
func StringSliceContains(slice []string, value string) bool {
	return StringSliceIndexOf(slice, value, 0) != IndexNotFound
}

// StringSliceIndexOf Finds the index of the given value in the slice.
func StringSliceIndexOf(slice []string, value string, startIndex int) int {
	if StringSliceIsEmpty(slice) {
		return IndexNotFound
	}
	if startIndex < 0 {
		startIndex = 0
	}
	length := len(slice)
	for i := startIndex; i < length; i++ {
		if value == slice[i] {
			return i
		}
	}
	return IndexNotFound
}

// StringSlicePrefixIndexOf Finds the index of the prefix of the given value in the prefix slice.
func StringSlicePrefixIndexOf(slice []string, value string, startIndex int) int {
	if StringSliceIsEmpty(slice) {
		return IndexNotFound
	}
	if startIndex < 0 {
		startIndex = 0
	}
	length := len(slice)
	for i := startIndex; i < length; i++ {
		prefix := slice[i]
		if strings.HasPrefix(value, prefix) {
			return i
		}
	}
	return IndexNotFound
}