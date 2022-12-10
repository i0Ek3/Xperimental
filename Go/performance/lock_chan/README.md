# Raw Lock vs Chan Lock

```Shell
➜  lock_chan git:(master) ✗ go version

go version go1.19.3 darwin/arm64

➜  lock_chan git:(master) ✗ go test -bench=.

goos: darwin
goarch: arm64
pkg: github.com/i0Ek3/go-learning/performance/lock_chan
BenchmarkChanMutex-8     	  451180	      2636 ns/op
BenchmarkCommonMutex-8   	  894684	      1344 ns/op
PASS
ok  	github.com/i0Ek3/go-learning/performance/lock_chan	2.438s
```
