package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
	//"github.com/i0Ek3/noerr"
)

func main() {
	start := time.Now()
	ch := make(chan string) // no buffered
	defa := "links.conf"

	for _, url := range os.Args[1:] {
		reading(url, defa, ch)
	}

	for range os.Args[1:] {
		file, err := os.OpenFile(
			"log.txt",
			os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
			0666,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed: %v\n", err)
		}
		defer file.Close()

		data := []byte(<-ch)
		n, err := file.Write(data)
		if err != nil || n == 0 {
			fmt.Fprintf(os.Stderr, "failed: %v\n", err)
		}
		//fmt.Fprintf(os.Stdout, <-ch)
		//fmt.Println(<-ch)
	}

	elapsed := time.Since(start).Seconds()
	fmt.Printf("%.2fs elapsed\n", elapsed)
}

func reading(url, defa string, ch chan<- string) {
	if len(url) == 0 {
		f, err := os.Open(defa)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed: %v\n", err)
		}
		f.Close()

		scanner := bufio.NewScanner(f)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			fmt.Println("-->", scanner.Text())
			go fetch(scanner.Text(), ch)
		}
	} else {
		go fetch(url, ch)
	}
}

func fetch(url string, ch chan<- string) {
	start := time.Now()

	prefix := []string{"http://", "https://"}
	if !strings.HasPrefix(url, prefix[0]) && !strings.HasPrefix(url, prefix[1]) {
		url = prefix[1] + url
	}

	resp, err := http.Get(url)
	//noerr.Xerr(err)
	if err != nil {
		//fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		ch <- fmt.Sprint(err)
		//os.Exit(1)
		return
	}

	//b, err := ioutil.ReadAll(resp.Body)
	//b, err := io.Copy(os.Stdout, resp.Body)
	b, err := io.Copy(ioutil.Discard, resp.Body)
	defer resp.Body.Close()
	if resp.Status != "200" {
		fmt.Println(resp.Status)
	}
	//noerr.Xerr(err)
	if err != nil {
		//fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		ch <- fmt.Sprint(err)
		//os.Exit(1)
		return
	}

	elapsed := time.Since(start).Seconds()

	//fmt.Printf("%s", b)
	ch <- fmt.Sprintf("%.2fs %7d %s", elapsed, b, url)
}
