package redis

import (
	"github.com/go-redis/redis"
)

var Client *redis.Client

func Connect() *redis.Client {
	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := Client.Ping().Result()
	if err != nil {
		panic(err)
	}
	return Client
}

func ConnectToRedis() *redis.Client {
	return Connect()
}
