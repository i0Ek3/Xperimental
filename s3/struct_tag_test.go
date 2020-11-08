package s3

import (
	"testing"
)

func TestTag(t *testing.T) {
	tt := TagType{true, "LV", 10}

	got := mtag(tt)
	want := true

	//asrt(t, got, want)
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
