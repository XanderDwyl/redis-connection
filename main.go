package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	redis "gopkg.in/redis.v5"
)

type smsp struct {
	Name string
	ID   int
}

func main() {
	log.Println("connecting to redis")
	redisc, err := GetRedisConnection()
	if err != nil {
		log.Println(fmt.Sprintf("Redis connection error: %s", err.Error()))
		return
	}

	log.Println(fmt.Sprintf("Redis connection success %s", redisc.Ping()))

	// test write to redis
	smspTest := smsp{"test", 1000}
	key := strconv.Itoa(smspTest.ID)

	// Set expiration to 2 hours
	redisSetErr := redisc.Set(key, key, 2*time.Hour).Err()
	if redisSetErr != nil {
		log.Println(fmt.Sprintf("Could not set redis with key"))
	}

	log.Println(
		fmt.Sprintf(
			"%s (%d) is successfully cached",
			smspTest.Name,
			smspTest.ID,
		))
}

// GetRedisConnection ...
func GetRedisConnection() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return client, nil
}
