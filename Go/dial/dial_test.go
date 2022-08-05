package dial

import (
	a "github.com/i0Ek3/asrt"
	"testing"
)

func TestDial(t *testing.T) {
	got := Dial()
	want := true

	a.Asrt(t, got, !want)
}
