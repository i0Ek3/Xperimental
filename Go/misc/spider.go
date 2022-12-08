package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	Data struct {
		Replies []struct {
			Content struct {
				Device  string        `json:"device"`
				JumpURL struct{}      `json:"jump_url"`
				MaxLine int64         `json:"max_line"`
				Members []interface{} `json:"members"`
				Message string        `json:"message"`
				Plat    int64         `json:"plat"`
			} `json:"content"`
			Replies []struct {
				Action  int64 `json:"action"`
				Assist  int64 `json:"assist"`
				Attr    int64 `json:"attr"`
				Content struct {
					Device  string   `json:"device"`
					JumpURL struct{} `json:"jump_url"`
					MaxLine int64    `json:"max_line"`
					Message string   `json:"message"`
					Plat    int64    `json:"plat"`
				} `json:"content"`
				Rcount  int64       `json:"rcount"`
				Replies interface{} `json:"replies"`
			} `json:"replies"`
		} `json:"replies"`
	} `json:"data"`
	Message string `json:"message"`
}

func main() {
	ch := make(chan struct{})
	client := &http.Client{}

	for i := 0; i < 100; i++ {
		go spider(client, ch)
	}

	for i := 0; i < 100; i++ {
		<-ch
	}
}

func spider(client *http.Client, ch chan struct{}) {
	apiURL := "https://api.bilibili.com/x/v2/reply/main?csrf=cfe5bd78ca65f2352e84d6e857519c1f&mode=3&next=1&oid=687054929&plat=1&type=1"
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("authority", "api.bilibili.com")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("referer", "https://www.bilibili.com/video/BV1PU4y1y7y9/?spm_id_from=333.788.recommend_more_video.-1&vd_source=f41e6a49a3bbca7077234de99c4c6b7d")
	req.Header.Set("accept-language", "en-US,en;q=0.9")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var result Response
	_ = json.Unmarshal(data, &result)

	for _, v := range result.Data.Replies {
		fmt.Println("[#1]:", v.Content.Message)
		for _, vv := range v.Replies {
			fmt.Println("  [#2]:", vv.Content.Message)
		}
	}

	if ch != nil {
		ch <- struct{}{}
	}
}
