package dal

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var (
	rdb *redis.Client
)
var ctx = context.Background()

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func SetKV(key, val string) {
	header := "[dal.SetKV]"
	err := rdb.Set(ctx, key, val, 0).Err()
	if err != nil {
		fmt.Printf("%v err: %v", header, err)
		panic(err)
	}
}

func GetKV(key string) string {
	header := "[dal.GetKV]"
	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		fmt.Printf("key (%v) does not exist\n", key)
	} else if err != nil {
		fmt.Printf("%v err: %v\n", header, err)
		panic(err)
	} else {
		fmt.Println(key, " : ", val)
	}
	return val
}
