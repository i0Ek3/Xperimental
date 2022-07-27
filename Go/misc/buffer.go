package main 

import ( 
        "fmt"
        "bytes"
)

func main() {
        var buffer bytes.Buffer
        for {
                if s, ok := getNextString(); ok {
                        buffer.WriteString(s)
                } else {
                        break 
                }
        }
        fmt.Print(buffer.String(), "\n")
}
