package main

import "fmt"

// Go中方法的定义：func (recv receiver_type) methodName(parameter_list) (return_value_list) { ... }
// 函数和方案的区别在于：函数将变量作为参数，而方法在变量上被调用。
// 函数：Function(recv)
// 方法：recv.Method()


type Ivector []int

type Ints struct {
        a, b int
}

type TZ int

//type A struct {
//        name string
//}

func main() {
        i1 := new(Ints)
        i1.a = 3
        i1.b = 4

        fmt.Printf("i1.a + i1.b = %d\n", i1.AddInts())
        fmt.Printf("i1.a + i1.b + 5 = %d\n", i1.AddParam(5))

        i2 := Ints{1, 3}
        fmt.Printf("i2.a + i2.b = %d\n", i2.AddInts())

        fmt.Printf("Ivector{1, 2, 3} = %d\n", Ivector{1, 2, 3}.Sum())
        
        var a TZ
        a.Print() // method value
        (*TZ).Print(&a) // method expression
    
}

// 方法可以访问私有字段，即其访问权限是比较高的。
func (a *TZ) Print() {
        //a.name = "john"
        fmt.Println("TZ")
}

func (it *Ints) AddInts() int {
        return it.a + it.b
}

func (it *Ints) AddParam(s int) int {
        return it.a + it.b + s
}

func (v Ivector) Sum() (s int) {
        for _, x := range v {
                s += x    
        }
        return
}


