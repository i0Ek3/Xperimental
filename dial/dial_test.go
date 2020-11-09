package dial

import (
	a "asrt"
	"testing"
)

func TestDial(t *testing.T) {
	got := Dial()
	want := true

	a.Asrt(t, got, want)
}
