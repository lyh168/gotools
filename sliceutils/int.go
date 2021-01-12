package sliceutils

import (
	"strconv"
)

// Int32SliceJoin concatenates the elements of its first argument to create a single string. The separator
// string sep is placed between elements in the resulting string.
func Int32SliceJoin(elems []int32, sep string) string {
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return strconv.Itoa(int(elems[0]))
	}
	var bs []byte

	bs = strconv.AppendInt(bs, int64(elems[0]), 10)
	for _, s := range elems[1:] {
		bs = append(bs, sep...)
		bs = strconv.AppendInt(bs, int64(s), 10)
	}
	return string(bs)
}
