package main 
import "fmt"

type test struct {
        i   int
        f   float32
        s   string
        t struct {
                t1, t2 string
        }
}

func main() {
        //tmp := new(test)
        //tmp := &test{1, 1.9, "test!"}
        //tmp.i = 1
        //tmp.f = 1.9
        //tmp.s = "test!"

        //fmt.Printf("tmp.i = %d\n", tmp.i)
        //fmt.Printf("tmp.f = %f\n", tmp.f)
        //fmt.Printf("tmp.s = %s\n", tmp.s)
        //fmt.Println(tmp)

        a := struct {
                Name string
                Age  int 
        }{
                Name: "andy",
                Age: 20,
        }
        fmt.Println(a)

        b := test{i: 1, f: 1.01, s: "1.01"}
        b.t.t1 = "This is test!"
        b.t.t2 = "Me too!"
        fmt.Println(b)
}
