package config

import (
	"github.com/gookit/ini/v2/dotenv"
	"log"
	"os"
)

func init() {
	err := dotenv.LoadExists("./", ".env")
	if err != nil {
		log.Panicln(err)
	}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}
