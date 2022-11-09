package main

import (
	"context"
	"fmt"

	"github.com/olivere/elastic/v7"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
}

func main() {
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"))
	if err != nil {
		panic(err)
	}
	fmt.Println("connect to es success")

	p := Person{Name: "i0Ek3", Age: 18, Married: false}
	put, err := client.Index().
		Index("user").
		BodyJson(p).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed user %s to index %s, type %s\n", put.Id, put.Index, put.Type)
}
