package emqx

import (
	"github.com/gofiber/fiber/v2"
	"os"
	"strings"
)

var allow = fiber.Map{
	"result": "allow",
}

var deny = fiber.Map{
	"result": "deny",
}

// "password": "${password}",
// "username": "${username}"
type authBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// HandleAuth http://fiber:648/emqx/auth
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
	return ctx.JSON(deny)
}

// "username": "${username}",
// "topic": "${topic}",
// "action": "${action}"
type authzBody struct {
	// 用户的 ID
	Username string `json:"username"`
	Topic    string `json:"topic"`
	Action   string `json:"action"`
}

// HandleAuthz http://fiber:648/emqx/authz
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
	if !strings.Contains(body.Topic, "#") && !strings.Contains(body.Topic, "+") {
		return ctx.JSON(allow)
	}
	return ctx.JSON(deny)
}
