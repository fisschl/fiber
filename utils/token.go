package utils

import "github.com/gofiber/fiber/v2"

func AuthToken(ctx *fiber.Ctx) string {
	token := ctx.Cookies("token")
	if token != "" {
		return token
	}
	token = ctx.Get("token")
	if token != "" {
		return token
	}
	return ctx.Query("token")
}
