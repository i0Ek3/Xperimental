package main 

import "fmt"

type In struct {
        a, b int
}

type Out struct {
        a int
        b, c float32
        int
        In
}

type Mix struct {
        In
        Out
}

func main() {
        
        var i In
        var o Out
        var m Mix

        i = In{1, 2}
        o = Out{1, 2.0, 3.0, 10, In{3, 4}}
        m = Mix{In{3, 4}, Out{1, 1.0, 1.1, 20, In{10, 11}}}

        fmt.Println(i, o, m)
        fmt.Printf("--------------------------\n")
        fmt.Println(i.a, i.b, o.a, o.b, o.c, o.int, o.In.a, o.In.b, m.In.a, m.In.b, m.Out.a, m.Out.b, m.Out.c, m.Out.int, m.Out.In.a, m.Out.In.b)
}

