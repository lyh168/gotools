package sliceutils

import (
	"testing"
)

func TestInt32SliceJoin(t *testing.T) {
	str := Int32SliceJoin([]int32{1, 2, 3, 4, 5, 6}, ",")
	if str == "1,2,3,4,5,6" {
		t.Log("success")
	} else {
		t.Error("failed")
	}
}
