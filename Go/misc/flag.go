package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
)

var (
	wb    string
	debug = flag.Bool("debug", false, "debug mode")

	usage = `Welcome to topic searcher:
./flag topic topicName
./flag -debug -wb=hn topic topicName`
)

func init() {
	flag.StringVar(&wb, "wb", "v2", "information source")

	flag.Parse()
}

func main() {
	if len(flag.Args()) < 1 {
		fmt.Println(usage)
		os.Exit(1)
	}
	cmd := flag.Args()[0]
	err := executeCmd(cmd, flag.Args()[1:], wb)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func executeCmd(cmd string, args []string, wb string) error {
	msg(fmt.Sprintf("command: %s", cmd))
	msg(fmt.Sprintf("args: %v", args))

	if cmd == "topic" {
		return Search(args, wb)
	} else {
		return fmt.Errorf("invalid command: %s", cmd)
	}
}

func Search(args []string, wb string) error {
	if len(args) == 0 {
		return errors.New("no enough argument")
	}

	topic := args[0]
	msg(fmt.Sprintf("topic: %s", topic))

	topics, err := search(topic, wb)
	if err != nil {
		return err
	}
	fmt.Println(strings.Join(topics, ", "))

	return nil
}

func search(topic string, wb string) ([]string, error) {
	type Info struct {
		Title string `json:"title"`
		Topic string `json:"topic"`
	}

	type result struct {
		Items []Info `json:"items"`
	}

	var url string
	if wb == "v2" {
		url = "https://www.v2ex.com/api/v2/topics"
	} else {
		url = "https://hacker-news.firebaseio.com/v0/topstories.json/item"
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		msg(fmt.Sprintf("%v", err))

		return nil, fmt.Errorf("%s no response", wb)
	}

	query := req.URL.Query()
	query.Set("q", topic)
	req.URL.RawQuery = query.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		msg(fmt.Sprintf("%v", err))

		return nil, fmt.Errorf("%s no response", wb)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("%s no response", wb)
	}

	res := result{}
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		msg(fmt.Sprintf("%v", err))

		return nil, fmt.Errorf("%s no response", wb)
	}

	topics := make([]string, 0)
	for _, r := range res.Items {
		if wb == "v2" {
			topics = append(topics, r.Topic)
		} else {
			topics = append(topics, r.Title)
		}
	}

	return topics, nil
}

func msg(s string) {
	if *debug {
		fmt.Printf("[debug] => %s\n", s)
	}
}
