package main

import (
	"fmt"
)

type T struct{}

func (t T) Hi() {
	fmt.Println("hi, I'm value")
}

func (t *T) hi() {
	fmt.Println("hi, I'm pointer")
}

func test1() {
	var t T
	t.Hi()
	fmt.Println("ok, t have method Hi()")
	t.hi()
	fmt.Println("ok, t also can invoke pointer type hi()")

	fmt.Println("----------------------------------")

	var pt *T
	pt.hi()
	fmt.Println("ok, cause of pt is a pointer")
	pt.Hi()
	fmt.Println("not ok, pt have no method Hi()")
}

func main() {
	test1()
}
