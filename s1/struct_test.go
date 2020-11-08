package s1

import "testing"

func TestStruct(t *testing.T) {
	b := test{i: 1, f: 1.01, s: "1.01"}

	got := mystruct(b)
	want := true

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
