package utils

import (
	"github.com/fisschl/fiber/config"
	"github.com/redis/go-redis/v9"
	"log"
)

var Rdb *redis.Client

func init() {
	opts, err := redis.ParseURL(config.GetEnv("REDIS_URL"))
	if err != nil {
		log.Panicln(err)
	}
	Rdb = redis.NewClient(opts)
}
