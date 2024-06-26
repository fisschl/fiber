package emqx

import (
	"github.com/fisschl/fiber/utils"
	"github.com/gofiber/fiber/v2"
	"os"
	"strings"
)

type authBody struct {
	// 用户的 ID
	Username string `json:"username"`
	// 用户的 Token
	Password string `json:"password"`
}

var allow = fiber.Map{
	"result": "allow",
}

var deny = fiber.Map{
	"result": "deny",
}

// HandleAuth http://fiber:648/emqx/auth
//
//	{
//	 "password": "${password}",
//	 "username": "${username}"
//	}
func HandleAuth(ctx *fiber.Ctx) error {
	var body authBody
	err := ctx.BodyParser(&body)
	if err != nil {
		return err
	}
	if body.Username == "default" && body.Password == os.Getenv("DEFAULT_PASSWORD") {
		return ctx.JSON(allow)
	}
	if body.Username == "public" && body.Password == "public" {
		return ctx.JSON(allow)
	}
	user, _ := utils.Rdb.HGet(ctx.Context(), body.Password, "user").Result()
	if user != "" && user == body.Username {
		return ctx.JSON(allow)
	}
	return ctx.JSON(deny)
}

type authzBody struct {
	// 用户的 ID
	Username string `json:"username"`
	Topic    string `json:"topic"`
	Action   string `json:"action"`
}

// HandleAuthz http://fiber:648/emqx/authz
//
//	{
//	 "username": "${username}",
//	 "topic": "${topic}",
//	 "action": "${action}"
//	}
func HandleAuthz(ctx *fiber.Ctx) error {
	var body authzBody
	err := ctx.BodyParser(&body)
	if err != nil {
		return err
	}
	if body.Action == "publish" {
		return ctx.JSON(allow)
	}
	if body.Username == "default" {
		return ctx.JSON(allow)
	}
	if strings.HasPrefix(body.Topic, "public") && !strings.Contains(body.Topic, "#") && !strings.Contains(body.Topic, "+") {
		return ctx.JSON(allow)
	}
	if body.Username == "public" {
		return ctx.JSON(deny)
	}
	if strings.HasPrefix(body.Topic, body.Username) {
		return ctx.JSON(allow)
	}
	return ctx.JSON(deny)
}
