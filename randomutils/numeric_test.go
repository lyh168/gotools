package randomutils

import (
	"strings"
	"testing"
)

func TestNumericSequence(t *testing.T) {
	for i := 1; i < 10; i++ {
		sequence := NumericSequence(i)
		if len(strings.TrimSpace(sequence)) != i {
			t.Errorf("expected numeric sequence length is %d, but result is %d", i, len(strings.TrimSpace(sequence)))
		} else {
			t.Logf("length %d numeric sequence is %s", i, sequence)
		}
	}

}
