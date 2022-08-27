package main 

import (
        "fmt"
        "bufio"
        "os"
        "compress/gzip" 
)
        
// read gzip package 

func main() {
        fileName := "file.gz" // you can touch one 
        var r *bufio.Reader
        f, err := os.Open(fileName)
        if err != nil {
                fmt.Fprintf(os.Stderr, "%v, Cannot open %s: error: %s\n", os.Args[0], fileName, err)
                os.Exit(1)
        }
        defer f.Close()
        gz, err := gzip.NewReader(f)
        if err != nil {
                r = bufio.NewReader(f)
        } else {
                r = bufio.NewReader(gz)                
        }
        
        for {
                line, err := r.ReadString('\n')
                if err != nil {
                        fmt.Println("Read file successful.")
                        os.Exit(0)
                }
                fmt.Println(line)
        }
}





