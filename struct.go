package main 
import "fmt"

type test struct {
        i   int
        f   float32
        s   string
}

func main() {
        //tmp := new(test)
        tmp := &test{1, 1.9, "test!"}
        tmp.i = 1
        tmp.f = 1.9
        tmp.s = "test!"

        fmt.Printf("tmp.i = %d\n", tmp.i)
        fmt.Printf("tmp.f = %f\n", tmp.f)
        fmt.Printf("tmp.s = %s\n", tmp.s)
        fmt.Println(tmp)
}
