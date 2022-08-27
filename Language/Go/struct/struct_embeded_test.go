package struct

import (
	"testing"

	a "asrt"
)

func TestEmbed(t *testing.T) {
	i := In{1, 2}
	o := Out{1, 2.0, 3.0, 10, In{3, 4}}
	m := Mix{In{3, 4}, Out{1, 1.0, 1.1, 20, In{10, 11}}}

	got := membed(i, o, m)
	want := false

	a.Asrt(t, got, want)
}
