package netUtils

import "testing"

func TestPickFreeAddr(t *testing.T) {
	l, err := PickFreePort()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(l)
}