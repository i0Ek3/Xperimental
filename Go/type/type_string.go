package main 

import (
        "fmt"
)

type T struct {
        a int
        b float32
        c string
}

func main() {
        //tn := new(T)
        //tn.a = 10
        //tn.b = 3.3345345
        //tn.c = "asb\tdee"
        //fmt.Printf("tn-v = %v\n", tn)
        //fmt.Println("tn =", tn)
        //fmt.Printf("tn-T = %T\n", tn)
        //fmt.Printf("tn-#v = %#v\n", tn)
        tn := &T{7, -2.35, "abc\tdef"}
        fmt.Printf("%v\n", tn)
}

func (t *T) String() string {
        return fmt.Sprintf("%d / %f / %q", t.a, t.b, t.c)
}
