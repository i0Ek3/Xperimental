package s2

import (
	"testing"
)

func TestEmbed(t *testing.T) {
	i := In{1, 2}
	o := Out{1, 2.0, 3.0, 10, In{3, 4}}
	m := Mix{In{3, 4}, Out{1, 1.0, 1.1, 20, In{10, 11}}}

	got := membed(i, o, m)
	want := false

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
