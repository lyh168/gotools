package sliceUtils

//Checks if an slice is nil or length equals 0
func StringSliceIsEmpty(slice []string) bool {
	return len(slice) <= 0
}

// Checks if the value is in the given slice.
func StringSliceContains(slice []string, value string) bool {
	return StringSliceIndexOf(slice, value, 0) != INDEX_NOT_FOUND
}

// Finds the index of the given value in the slice.
func StringSliceIndexOf(slice []string, value string, startIndex int) int {
	if StringSliceIsEmpty(slice) {
		return INDEX_NOT_FOUND
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
	return INDEX_NOT_FOUND
}
