// sudo apt-get update
// sudo apt-get install redis
// sudo service redis-server start

package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis"
)

var (
	address  = "localhost:6379"
	password = ""
)

type RedisDatabase struct {
	Client *redis.Client
}

func NewRedis() (*RedisDatabase, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       0, // use default DB
	})

	if _, err := client.Ping(context.TODO()).Result(); err != nil {
		return nil, err
	}

	return &RedisDatabase{
		Client: client,
	}, nil
}

func main() {
	redisdb, err := NewRedis()
	if err != nil {
		panic(err)
	}

	key := "key1"
	value := "value1"
	// value := true

	// write
	if err := redisdb.Client.Set(context.TODO(), key, value, 0).Err(); err != nil {
		panic(err)
	}

	// read
	val, err := redisdb.Client.Get(context.TODO(), key).Result() // => GET key
	// val, err := redisdb.Client.Get(context.TODO(), key).Bool() // => GET key
	if err == redis.Nil {
		fmt.Println("not exist")
	} else if err != nil {
		panic(err)
	}
	fmt.Println(key, val)
}
