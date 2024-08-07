package utils

import (
	"github.com/gofiber/fiber/v2/log"
	"io"
	"os"
)

func GetEnv(key string) string {
	return os.Getenv(key)
}

func Close(item io.Closer) {
	err := item.Close()
	if err != nil {
		log.Error(err)
	}
}
