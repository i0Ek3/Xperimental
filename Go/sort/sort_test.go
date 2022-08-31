package msort

import (
	"testing"

	a "github.com/i0Ek3/asrt"
)

var (
	i IntArray
)

const (
	n1 = 1
	n2 = 2
)

func TestSwap(t *testing.T) {
	got := i.Swap(n1, n2)
	want := true

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	//a.Asrt(t, got, want)
}

func TestLess(t *testing.T) {
	got := i.Less(n1, n2)
	want := true

	a.Asrt(t, got, want)
}
