package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v9"
)

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	err := rdb.Set(ctx, "key", "val", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("non-exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}

    err = rdb.Publish(ctx, "chan1", "payload").Err()
    if err != nil {
        panic(err)
    }

    pubsub := rdb.Subscribe(ctx, "chan1")
    defer pubsub.Close()

    ch := pubsub.Channel()
    for msg := range ch {
        fmt.Println(msg.Channel, msg.Payload)
    }
}
