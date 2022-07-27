package main

import "fmt"

// go语言中，通过接口来实现面向对象的种种
// 其中，接口定义了一组方法，它们是抽象的，且接口中不能包含变量
// 其定义如下：

// type Namer interface {
// 		// Namer 是接口类型，接口的名字由方法名加er后缀组成，例如Printer，Reader等
//		// 当后缀er不合适时，接口通常以able结尾，或者以I开头
//		// Go 语言中的接口都很简短，通常它们会包含 0 个、最多 3 个方法
//
// 		Method1(param_list) return_type
//		Method2(param_lisr) return_type
// 		...
// }

type Shaper interface {
		Area() float32
}

type Square struct {
		side float32
}

func (sq * Square) Area() float32 {
		return sq.side * sq.side
}

type Rectangle struct {
		length, width float32
}

func (r Rectangle) Area() float32 {
		return r.length * r.width
}

func main() {
		r := Rectangle{5, 3}
		q := &Square{4}
		shapes := []Shaper{r, q}
		fmt.Println("Looping through shapes for area...")
		for n, _ := range shapes {
				fmt.Println("Shape details: ", shapes[n])
				fmt.Println("Area of this shape is: ", shapes[n].Area())
		}

}
