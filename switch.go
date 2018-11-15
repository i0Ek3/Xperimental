package main

import "fmt"

//
//switch i {
//        case 0: fallthrough
//        case 1:
//                do()
//        default:
//                to()
//}
// 

func main() {
        var num int = 100
        switch num {
        case 98, 99:
                fmt.Println("98\n00")
        case 100:
                fmt.Println("100\n")
        default:
                fmt.Println("none\n")
        }

        var i int = 3
        switch {
                case i < 0:
                        fmt.Println("do1()")
                case i == 0:
                        fmt.Println("do2()")
                case i > 5:
                        fmt.Println("do3()")
                default:
                        fmt.Println("Bye~")
        }
}
