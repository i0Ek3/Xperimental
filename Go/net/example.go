package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "https://github.com/i0Ek3"

	data := []byte(`{
		"name": "i0Ek3",
		"job": "freelancer",
	}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("resp status:", resp.Status)
	fmt.Println("resp header:", resp.Header)
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("resp body:", string(body))
}
