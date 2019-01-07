# Go

There are awesome go samples.

Also you can visit official go sites to know more.

> - [Go语言英文网](http://docs.studygolang.com/)
> - [Go语言中文网](https://studygolang.com/)




## Goenv

Download corresponding package from official site then run as follows: 

```Shell
$ sudo tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz
$ export PATH=$PATH:/usr/local/go/bin` and add this line in your /etc/profile.

```

macOS can use homebrew install go with command: `brew install go`.


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

package main // which belongs to.

import "fmt" // package implementing formatted I/O.
//import fm "fmt" // alias.
import os // imported and not used: os, cause of there is no need for `os`.


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

- 不要试图make()一个结构体，这样会引发一个编译错误
- 不要试图new()一个映射，这样会引发运行时错误
- 使用锁的情景：
    - 访问共享数据结构中的缓存信息
    - 保存应用程序上下文和状态信息数据
- 使用通道的情景：
    - 与异步操作的结果进行交互
    - 分发任务
    - 传递数据所有权



