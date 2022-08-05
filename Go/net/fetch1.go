package main

import (
	"fmt"
	//"io/ioutil"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		fetch := func(url string) {
			resp, err := http.Get(url)
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
				os.Exit(1)
			}
			code := resp.Status
			fmt.Printf("http status code: %s\n", code)
			//b, err := ioutil.ReadAll(resp.Body)
			b, err := io.Copy(os.Stdout, resp.Body) // recommend
			resp.Body.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
				os.Exit(1)
			}
			fmt.Printf("%s", b)
		}

		if strings.HasPrefix(url, "http://") {
			fetch(url)
		} else {
			url = "http://" + url
			fetch(url)
		}
	}
}
