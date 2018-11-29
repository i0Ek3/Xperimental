package main 

import (
        "fmt"
        "os"
        "io"
        "net/http"
        //"strings"
) 

// fetch一个网站

func main() {
        for _, url := range os.Args[1:] {
                resp, err := http.Get(url) // strings.HasPrefix()用来拉取那些不支持http协议的网站
                if err != nil {
                        fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
                        os.Exit(1) // 用resp.Status来获取状态码
                }
                b, err := io.Copy(os.Stdout, resp.Body)
                resp.Body.Close()
                if err != nil {
                        fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
                }
                fmt.Printf("%s", b)
        }
}
