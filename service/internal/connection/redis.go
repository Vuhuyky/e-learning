package connection

import (
	"log"

	"github.com/redis/go-redis/v9"
)

func connectRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr: conn.Redis,
	})
	log.Println("redis connected")
}
