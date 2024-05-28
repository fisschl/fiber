package utils

import (
	"github.com/gookit/ini/v2/dotenv"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
)

var Rdb *redis.Client

func init() {
	err := dotenv.LoadExists("./", ".env")
	if err != nil {
		log.Panicln(err)
	}
	opts, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		log.Panicln(err)
	}
	Rdb = redis.NewClient(opts)
}
