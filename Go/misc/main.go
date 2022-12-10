package main

import (
    "fmt"
)

func main() {
    test6()
}

func test1() {
    var s fmt.Stringer
    //s = "s"
    fmt.Println(s)
}

func test2() {
    const X = 7.0
    var x interface{} = X

    if y, ok := x.(int); ok {
        fmt.Println(y)
    } else {
        fmt.Println(int(y))
    }
}

func test3() {
    s := []int{5}
	s = append(s, 6)
	s = append(s, 7)
	x := append(s, 11)
	fmt.Println(s, x)
	fmt.Println("------ x was overrided ------")
	y := append(s, 12)
	fmt.Println(s, x, y)
}

func test4() {
    c := make(chan int, 1)
	for done := false; !done; {
		select {
		default:
			fmt.Println(1)
		case <-c:
			fmt.Println(2)
			c = nil
		case c <- 1:
			fmt.Println(3)
		}
	}
}

func test5() {
    s := []string{"A", "B", "C"}
    t := s[:1]
    fmt.Println(&s[0] == &t[0])

    u := append(s[:1], s[2:]...)
    fmt.Println(&s[0] == &u[0])
}

type Test6 struct{}

func test6() {
    t := &Test6{}
    defer t.Add(1).Add(2)
    t.Add(3)
}

func (t *Test6) Add(x int) *Test6 {
    fmt.Println(x)
    return &Test6{}
}
