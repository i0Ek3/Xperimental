# Go Generic

## Build

Best way to run go2 code is build go2 from source code, and run command `go tool go2go run xxx.go2`. 

.go is different with .go2, cause of .go2 indicates the go2 version, you can use command `go tool go2go translate xxx.go2` to tanslate it to .go.

If you want use .go file to exprienced Go generic, please run command `go run -gcflags=-G=3 xxx.go`.

## Playground

[https://go2goplay.golang.org](https://go2goplay.golang.org).

## Refs

- [https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#generic-types](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#generic-types)
- [https://gocn.vip/topics/12460](https://gocn.vip/topics/12460)
- [https://colobu.com/2020/06/18/run-local-go-generic-files/](https://colobu.com/2020/06/18/run-local-go-generic-files/)
