package main 

import (
        "fmt"
)

//
// 在接口上调用方法时，必须有和方法定义时相同的接收者类型或者是可以从具体类型 P 直接可以辨识的：
//
//  指针方法可以通过指针调用
//  值方法可以通过值调用
//  接收者是值的方法可以通过指针调用，因为指针会首先被解引用
//  接收者是指针的方法不可以通过值调用，因为存储在接口中的值没有地址
//  将一个值赋值给一个接口时，编译器会确保所有可能的接口方法都可以在此值上被调用，因此不正确的赋值在编译期就会失败。
//
//
// Go 语言规范定义了接口方法集的调用规则：
//
//  类型 *T 的可调用方法集包含接受者为 *T 或 T 的所有方法集
//  类型 T 的可调用方法集包含接受者为 T 的所有方法
//  类型 T 的可调用方法集不包含接受者为 *T 的方法
//

type List []int

func (l List) Len() int {
        return len(l)
}

func (l *List) Append(val int) {
        *l = append(*l, val)
}


type Appender interface {
        Append(int)
}

type Lener interface {
        Len() int
}

func CountInto(a Appender, start, end int) {
        for i := start; i <= end; i++ {
                a.Append(i)
        }
}

func LongEnough(l Lener) bool {
        return l.Len() * 10 > 42
}

//
// # command-line-arguments
// ./methodset.go:32:1: missing return at end of function
// ./methodset.go:60:18: cannot use plst (type *List) as type Appender in argument to CountInto:
//        *List does not implement Appender (wrong type for Append method)
//               have Append(int) int
//                want Append(int)
//


func main() {
        var lst List
        if LongEnough(lst) {
                fmt.Printf("- lst long enough.\n")
        }

        plst := new(List)
        CountInto(plst, 1, 10)
        if LongEnough(plst) {
                fmt.Printf("- plst long enough.\n")
        }
}




