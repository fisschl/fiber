package emqx

import (
	"github.com/fisschl/fiber/utils"
	"github.com/gofiber/fiber/v2"
)

type authBody struct {
	// 用户的 ID
	Username string `json:"username"`
	// 用户的 Token
	Password string `json:"password"`
}

// http://fiber:648/emqx/auth
//
//	{
//	 "password": "${password}",
//	 "username": "${username}"
//	}
func authHandler(ctx *fiber.Ctx) error {
	var body authBody
	err := ctx.BodyParser(&body)
	if err != nil {
		return err
	}
	user, err := utils.Rdb.HGet(ctx.Context(), body.Password, "user").Result()
	if err != nil || user == "" {
		return ctx.SendStatus(fiber.StatusForbidden)
	}
	if user != body.Username {
		return ctx.SendStatus(fiber.StatusForbidden)
	}
	return ctx.JSON(fiber.Map{
		"result": "allow",
	})
}

type authzBody struct {
	// 用户的 ID
	Username string `json:"username"`
	Topic    string `json:"topic"`
	Action   string `json:"action"`
}

// http://fiber:648/emqx/authz
//
//	{
//	 "username": "${username}",
//	 "topic": "${topic}",
//	 "action": "${action}"
//	}
func authzHandler(ctx *fiber.Ctx) error {
	var body authzBody
	err := ctx.BodyParser(&body)
	if err != nil {
		return err
	}
	if body.Action == "publish" {
		return ctx.JSON(fiber.Map{
			"result": "allow",
		})
	}
	if body.Username != body.Topic {
		return ctx.SendStatus(fiber.StatusForbidden)
	}
	return ctx.JSON(fiber.Map{
		"result": "allow",
	})
}

func RegisterRouter(router fiber.Router) {
	router.Post("/emqx/auth", authHandler)
	router.Post("/emqx/authz", authzHandler)
}
