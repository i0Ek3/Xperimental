package mstring

import (
	"testing"

	a "github.com/i0Ek3/asrt"
)

func TestString(t *testing.T) {
	oldstr := "old"
	newstr := "OLD"

	got := mstring(oldstr, newstr)
	want := true

	a.Asrt(t, got, want)
}
