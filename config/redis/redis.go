package redis

import (
	"fmt"

	"github.com/go-redis/redis"
)

func Connect(addr string, password string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
	})

	pingRes, _ := client.Ping().Result()
	if pingRes != "PONG" {
		fmt.Println("Failed to connect redis.")
	} else if pingRes == "PONG" {
		fmt.Println("The Redis connection is successful: ", pingRes)
	}
	return client
}
