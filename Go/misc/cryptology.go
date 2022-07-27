package main

import (
        "fmt"
        "io"
        "log"
        "crypto/sha1"
        "crypto/md5"
)

// Of course Google add more cryptology packages into Go languages to prevent attacks from hackers. 
// Such as: hash, crypto and so on...
// Now let we write some stuffs to validate it.

type Hash interface {
        io.Writer
        Sum(b []byte) []byte
        Reset()
        Size() int
        BlockSize() int
}


func hasher() {
        hasher := sha1.New() // create a new object hash.Hash
        io.WriteString(hasher, "test")
        b := []byte{}
        fmt.Printf("hex: %x\n", hasher.Sum(b))
        fmt.Printf("dec: %d\n", hasher.Sum(b))

        hasher.Reset()
        data := []byte("This sentence will be crypted!")
        n, err := hasher.Write(data)
        if n != len(data) || err != nil {
                log.Printf("Hash write error: %v / %v", n, err)
        }
        checksum := hasher.Sum(b)
        fmt.Printf("checksum: %x\n", checksum)
}

func md5str() {
        md5str := md5.New()
        io.WriteString(md5str, "test")
        b := []byte{}
        fmt.Printf("hex: %x\n", md5str.Sum(b))
        fmt.Printf("dec: %d\n", md5str.Sum(b))

        md5str.Reset()
        data := []byte("This sentence will be crypted!")
        n, err := md5str.Write(data)
        if n != len(data) || err != nil {
                log.Printf("Hash write error: %v / %v", n, err)
        }
        checksum := md5str.Sum(b)
        fmt.Printf("checksum: %x\n", checksum)
    

}


func main() {
        fmt.Printf("------------------hasher()-------------------\n")
        hasher()
        fmt.Printf("------------------md5str()-------------------\n")
        md5str()
}
