# Go Learing

Go examples for myself to learn Go.

## Install Goenv

Download corresponding package from official site then run as follows command: 

```Shell
$ sudo tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz
$ export PATH=$PATH:/usr/local/go/bin` and add this line in your /etc/profile.

```

If you use macOS, just run command: `brew install go`.


## Build & Run

```Go
$ go build source.go  // compile and install packages
$ go run source.go    // run the go file
$ go install packages // same with go build but without compile

```


## Sample

```Go
//go.go
//there is no ; as end of every code sentence, cause of this task did by compiler automatically.

package main // which belongs to

import "fmt" // package implementing formatted I/O
//import fm "fmt" // alias
//import . "fmt" // not recommend
import os // imported and not used: os, cause of there is no need for `os`
//import _ "pkg" // only use init() in pkg
//import "abc/def" // can use all pkgs under of def/


//
// func functionName(/*parameter_list*/ param1 type1, param2 type2, ...) (/*return_value_list*/ ret1 type1, ret2 type2, ...) {
//       ...
// }
//


// Calling show us something can be called by external.
func Calling() {
        // this function can be called by external packages, because this function start with upper letter.
        // that's mean function start with lower letter cannot called by external packages.
}


// if there is no main(), then will trigger build error.
func main() { // function, if there is init(), then first run func is init().
        fm.Println("Hello, Go!"); // output
        fmt.Printf("Καλημέρα κόσμε; or こんにちは 世界\n") // we can use international characters in go.
}

```

## Note

- [go_tests#ysk](https://github.com/i0Ek3/go_tests#ysk)


## Links you want to know

- [Go语言英文网](http://docs.studygolang.com/)
- [Go语言中文网](https://studygolang.com/)
- [https://blog.golang.org](https://blog.golang.org)
- [https://talks.golang.org](https://talks.golang.org)
- [https://play.golang.org](https://play.golang.org)

