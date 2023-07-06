package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"testing"
)

func Test_conn(t *testing.T){
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(context.Background(), "key", "value001", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(context.Background(), "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
}
