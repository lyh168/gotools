package sliceutils

import "testing"

func TestStringSlicePrefixIndexOf(t *testing.T) {
	idx := StringSlicePrefixIndexOf([]string{"aaaaa", "bbbbb", "ccccc","ddddd"}, "ccccc_123456", 0)
	if idx != 2 {
		t.Errorf("expect is 2 but actul is %d", idx)
	}

	idx = StringSlicePrefixIndexOf([]string{"apple", "banan", "orange","pair"}, "pairsweet", 2)
	if idx != 3 {
		t.Errorf("expect is 3 but actul is %d", idx)
	}
}