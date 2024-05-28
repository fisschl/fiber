package oss

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRouter(router fiber.Router) {
	router.Get("/oss/sts", stsHandler)
	router.Get("/oss/download", downloadHandler)
}
