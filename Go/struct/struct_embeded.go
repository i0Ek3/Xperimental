package struct

import "fmt"

var (
	i = In{1, 2}
	o = Out{1, 2.0, 3.0, 10, In{3, 4}}
	m = Mix{In{3, 4}, Out{1, 1.0, 1.1, 20, In{10, 11}}}
)

type In struct {
	a int
	b int
}

type Out struct {
	a int
	b float32
	c float32
	int
	In
}

type Mix struct {
	In
	Out
}

func membed(i In, o Out, m Mix) bool {
	s1 := fmt.Sprint(i, o, m)
	s2 := fmt.Sprint(i.a, i.b, o.a, o.b, o.c, o.int, o.In.a, o.In.b, m.In.a, m.In.b, m.Out.a, m.Out.b, m.Out.c, m.Out.int, m.Out.In.a, m.Out.In.b)
	if s1 == s2 {
		return true
	}
	return false
}
