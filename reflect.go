package main

import (
        "fmt"
        "reflect"
)

//
// 反射是用程序检查其所拥有的结构(不如类型)的一种能力, 这是元编程的一种形式。
// 反射可以在运行时检查类型和变量，例如它的大小、方法和动态的调用这些方法。
// 反射是通过检查一个接口的值，变量首先被转换成空接口。
// 
// func TypeOf(i interface{}) Type
// func ValueOf(i interface{}) Value
// 

func main() {
        var x1 float64 = 3.4
        fmt.Println("x1 = ", x1)
	    fmt.Println("type:", reflect.TypeOf(x1))
	    v1 := reflect.ValueOf(x1)
	    fmt.Println("value:", v1)
	    fmt.Println("type:", v1.Type())
	    fmt.Println("kind:", v1.Kind())
	    fmt.Println("value:", v1.Float())
	    fmt.Println(v1.Interface())
	    fmt.Printf("value is %5.2e\n", v1.Interface())
	    y1 := v1.Interface().(float64)
	    fmt.Println(y1)

        fmt.Printf("\n----------------------------------\n")

        var x2 float64 = 3.4
        fmt.Println("x2 = ", x2)
	    v2 := reflect.ValueOf(x2)
	    // setting a value:
	    // v.SetFloat(3.1415) // Error: will panic: reflect.Value.SetFloat using unaddressable value
	    fmt.Println("settability of v2:", v2.CanSet())
	    v2 = reflect.ValueOf(&x2) // Note: take the address of x.
	    fmt.Println("type of v2:", v2.Type())
	    fmt.Println("settability of v2:", v2.CanSet())
	    v2 = v2.Elem()
	    fmt.Println("The Elem of v2 is: ", v2)
	    fmt.Println("settability of v2:", v2.CanSet())
	    v2.SetFloat(3.1415) // this works!
	    fmt.Println(v2.Interface())
	    fmt.Println(v2)
}





