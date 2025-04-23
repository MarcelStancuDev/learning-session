package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

var RDB *redis.Client
var CTx = context.Background()

func InitRedis() {
	RBD = redis.NewClient(&redis.Options{
		Addr:	 "localhost:6379",
	})
	_, err := RDB.Ping(CTx).Result()
	if err != nil {
		log.Fatal("Redis error:", err)
	}
	log.Println("Connected to Redis")
}