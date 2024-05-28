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

func UserId(ctx *fiber.Ctx) string {
	token := AuthToken(ctx)
	if token == "" {
		return ""
	}
	user := Rdb.HGet(ctx.Context(), token, "user").String()
	return user
}
