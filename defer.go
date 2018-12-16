package main

import ( 
        "fmt"
        "os"
        "io/ioutil"
)

func main() {
        // defer存在的意义在于释放某些已分配的资源,类似于c++中的auto关键字
        fmt.Println("-------withDefer--------\n")
        first()
        fmt.Println("-------withNoDefer--------\n")
        withNoDefer()
        fmt.Println("-------new one--------\n")
        data("newone")
}

func withNoDefer() {
        fmt.Println("I am the first.\n")
        second()
        fmt.Println("I am the third.\n")
}

func first() {
        fmt.Println("I am the first.\n")
        defer second()
        fmt.Println("I am the third.\n")
}

func second() {
        fmt.Println("I am the second.\n")
}

func data(name string) string {
        f, _ := os.OpenFile(name, os.O_RDONLY, 0)
        defer f.Close()
        contents, _ := ioutil.ReadAll(f)
        return string(contents)
}
