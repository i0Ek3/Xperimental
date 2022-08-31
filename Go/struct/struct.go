package struct

import "fmt"

type test struct {
	i int
	f float32
	s string

	t struct {
		t1 string
		t2 string
	}
}

func mystruct(b test) bool {
	//tmp := new(test)
	//tmp := &test{1, 1.9, "test!"}
	//tmp.i = 1
	//tmp.f = 1.9
	//tmp.s = "test!"

	//fmt.Printf("tmp.i = %d\n", tmp.i)
	//fmt.Printf("tmp.f = %f\n", tmp.f)
	//fmt.Printf("tmp.s = %s\n", tmp.s)
	//fmt.Println(tmp)

	/*
		a := struct {
			Name string
			Age  int
		}{
			Name: "andy",
			Age:  20,
		}
		sa := fmt.Sprint(a)
		if sa != "" {
			return true
		} else {
			return false
		}
	*/

	b = test{i: 1, f: 1.01, s: "1.01"}
	b.t.t1 = "This is test!"
	b.t.t2 = "Me too!"
	sb := fmt.Sprint(b)
	if sb != "" {
		return true
	} else {
		return false
	}
}