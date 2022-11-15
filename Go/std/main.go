package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

func TestNewWriter() {
	w := bufio.NewWriter(os.Stdout)
	fmt.Fprint(w, "Hello, ")
	fmt.Fprint(w, "world!")
	for _, i := range []int64{1, 2, 3, 4} {
		b := w.AvailableBuffer()
		b = strconv.AppendInt(b, i, 10)
		b = append(b, ' ')
		w.Write(b)
	}
	w.Flush()
}

func TestBuffer() {
	var b bytes.Buffer
	b.Write([]byte("halo "))
	b.WriteString("Go ")
	fmt.Fprintf(&b, "language! ")
	b.WriteTo(os.Stdout)

	buf := bytes.NewBufferString("R29waGVycyBydWxlIQ==")
	dec := base64.NewDecoder(base64.StdEncoding, buf)
	io.Copy(os.Stdout, dec)
}

func TestWithCancel() {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

const shortDur = 1 * time.Millisecond

func TestWithDeadline() {
	d := time.Now().Add(shortDur)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

func TestWithTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), shortDur)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

func TestWithValue() {
	type CtxKey string
	f := func(ctx context.Context, k CtxKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found:", v)
			return
		}
		fmt.Println("not found:", k)
	}

	k := CtxKey("language")
	ctx := context.WithValue(context.Background(), k, "Go")

	f(ctx, k)
	f(ctx, CtxKey("color"))
}

func TestGroup() {
	g := new(errgroup.Group)
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
	}
	for _, u := range urls {
		url := u
		g.Go(func() error {
			resp, err := http.Get(url)
			if err == nil {
				resp.Body.Close()
			}
			return err
		})
	}
	if err := g.Wait(); err == nil {
		fmt.Println("ok!")
	}
}

func TestFmt() {
	const name, id = "i0Ek3", 18
	err := fmt.Errorf("-> %q (id %d) not found", name, id)
	fmt.Println(err.Error())

	n, err := fmt.Fprint(os.Stdout, name, " is ", id, " years old.\n")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprint: %v\n", err)
	}
	fmt.Print(n, " bytes written.\n")
}

func TestPool() {
	var w io.Writer = os.Stdout
	key := "path"
	val := "/search?q=flowers"
	var bufPool = sync.Pool{
		New: func() any {
			return new(bytes.Buffer)
		},
	}
	now := time.Unix(1136214245, 0)

	b := bufPool.Get().(*bytes.Buffer)
	b.Reset()
	b.WriteString(now.UTC().Format(time.RFC3339))
	b.WriteByte(' ')
	b.WriteString(key)
	b.WriteByte('=')
	b.WriteString(val)
	w.Write(b.Bytes())
	bufPool.Put(b)
}

type httpPkg struct{}

func (httpPkg) Get(url string) {}

var http httpPkg

func TestWaitGroup() {
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.example.com/",
	}
	for _, u := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			http.Get(url)
		}(u)
	}
	wg.Wait()
}

func statusUpdate() string { return "" }
func TestTick() {
	t := time.Tick(3 * time.Second)
	for next := range t {
		fmt.Printf("%v %s\n", next, statusUpdate())
	}
}

func handle(int) {}
func TestAfterSleep() {
	var c chan int
	select {
	case m := <-c:
		handle(m)
	case <-time.After(1 * time.Second):
		time.Sleep(100 * time.Microsecond)
		fmt.Println("timed out")
	}
}

func TestParseDuration() {
	h, _ := time.ParseDuration("10h")
	c, _ := time.ParseDuration("1h10m10s")
	m, _ := time.ParseDuration("1µs")
	n, _ := time.ParseDuration("1us")
	fmt.Println(h, c, m, n)
}

func TestNewTicker() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	done := make(chan bool)
	go func() {
		time.Sleep(3 * time.Second)
		done <- true
	}()
	for {
		select {
		case <-done:
			fmt.Println("done")
			return
		case t := <-ticker.C:
			fmt.Println("current time:", t)
		}
	}
}

func TestAdd() {
	start := time.Date(2009, 1, 1, 12, 0, 0, 0, time.UTC)
	afterTenS := start.Add(10 * time.Second)
	afterTenH := start.Add(10 * time.Hour)
	fmt.Printf("start = %v\n", start)
	fmt.Printf("after 10s = %v\n", afterTenS)
	fmt.Printf("after 10h = %v\n", afterTenH)
}

func TestUnix() {
	unixTime := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println(unixTime.Unix())
	t := time.Unix(unixTime.Unix(), 0).UTC()
	fmt.Println(t)
}

func main() {
	TestNewWriter()
	TestBuffer()
	TestWithCancel()
	TestWithDeadline()
	TestWithTimeout()
	TestWithValue()
	TestGroup()
	TestFmt()
	TestPool()
	TestWaitGroup()
	TestTick()
	TestAfterSleep()
	TestParseDuration()
	TestNewTicker()
	TestAdd()
	TestUnix()
}
