package b

import "testing"

func TestB(t *testing.T) {
	if B() {
		t.Log(true)
	} else {
		t.Error(false)
	}
}
