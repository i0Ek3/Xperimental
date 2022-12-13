# Raw Lock vs Chan Lock

```Shell
➜  lock_chan git:(master) go version
go version go1.19.3 darwin/arm64
➜  lock_chan git:(master) go test -bench=..
goos: darwin
goarch: arm64
pkg: github.com/i0Ek3/Xperimental/performance/lock_chan
BenchmarkChanMutex-8     	  450782	      2632 ns/op
BenchmarkCommonMutex-8   	  894043	      1345 ns/op
PASS
ok  	github.com/i0Ek3/Xperimental/performance/lock_chan	2.436s
```
