package main

import (
        "fmt"
        "io/ioutil"
        "os"
)

func main() {
        read_to_string()
        read_by_coloum()
}

func read_to_string() {
        inputFile := "file.txt"
        outputFile := "copy.txt"
        buf, err := ioutil.ReadFile(inputFile)
        if err != nil {
                fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
        }
        // buf := make([]byte, 1024)
        // n, err := inputReader.Read(buf)
        // if (n == 0) { break }

        fmt.Printf("%s\n", string(buf))
        err = ioutil.WriteFile(outputFile, buf, 0644)
        if err != nil {
                panic(err.Error())
        }
}

func read_by_coloum() {
        file, err := os.Open("file.txt")
        if err != nil {
                panic(err)
        }

        defer file.Close()

        var c1, c2, c3 []string
        for {
                var v1, v2, v3 string
                _, err := fmt.Fscanln(file, &v1, &v2, &v3)
                if err != nil {
                        break 
                }
                c1 = append(c1, v1)
                c2 = append(c2, v2)
                c3 = append(c3, v3)
        }
        fmt.Println(c1)
        fmt.Println(c2)
        fmt.Println(c3)
}


