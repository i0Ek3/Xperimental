package main 

import (
        "os"
        "fmt"
        "flag"
        "strings"
)

func use_os_to_read() {
        who := "mesa "
        if len(os.Args) > 1 {
                who += strings.Join(os.Args[1:], " ")
        }
        fmt.Println("Hello", who)
}

var NewLine = flag.Bool("n", false, "print newline")

const (
      Space   = " "
      newLine = "\n"
)

func use_flag_to_read() {
        flag.PrintDefaults()
        flag.Parse()
        var s string = ""
        for i := 0; i < flag.NArg(); i++ {
                if i > 0 {
                        s += " "
                        if *NewLine {
                                s += newLine
                        }
                }
                s += flag.Arg(i)
        }
        os.Stdout.WriteString(s)
}

func main() {
        //use_os_to_read()
        use_flag_to_read()
}



