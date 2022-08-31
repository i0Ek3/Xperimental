package struct

import (
	"testing"

	a "asrt"
)

func TestStruct(t *testing.T) {
	b := test{i: 1, f: 1.01, s: "1.01"}

	got := mystruct(b)
	want := true

	a.Asrt(t, got, want)
}
