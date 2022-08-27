package struct

import (
	"testing"

	a "asrt"
)

func TestTag(t *testing.T) {
	tt := TagType{true, "LV", 10}

	got := mtag(tt)
	want := true

	a.Asrt(t, got, want)
}
