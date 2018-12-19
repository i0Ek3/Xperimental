package main

import (
        "fmt"
        "os"
        "log"
        "encoding/json"
)

type Address struct {
        Type        string
        City        string
        Country     string 
}

type VCard struct {
        FirstName   string
        LastName    string
        Address     []*Address
        Remark      string 
}

func main() {
        pa := &Address{"private", "Aartselaar", "Belgium"}
        wa := &Address{"work", "Beijing", "China"}
        vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
        js, _ := json.Marshal(vc)
        fmt.Printf("JSON format: %s", js)
        file, _ := os.OpenFile("vcard.json", os.O_CREATE|os.O_WRONLY, 0666)
        defer file.Close()
        //end := json.NewDecoder(file) //unserialization
        //err := json.Decode(vc)
        enc := json.NewEncoder(file)
        err := enc.Encode(vc)
        if err != nil {
                log.Println("Error in encoding json.")
        }
}
