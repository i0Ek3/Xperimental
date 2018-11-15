package main 

import (
        "fmt"
        "regexp"
        "strconv"
)

func main() {
        searchIn := "hello.0hEllo.HELLO.2GHelloHeLLO"
        pat := "[a-zA-Z]+.[0-9]"

        f := func(s string) string {
                v, _ := strconv.ParseFloat(s, 32)
                return strconv.FormatFloat(v, 'f', 2, 32)
        }

        if ok, _ := regexp.MatchString(pat, searchIn); ok {
                fmt.Println("Match Found!")
        }

        re, _ := regexp.Compile(pat)
        str := re.ReplaceAllString(searchIn, "+")
        fmt.Println(str)

        str2 := re.ReplaceAllStringFunc(searchIn, f)
        fmt.Println(str2)

}
