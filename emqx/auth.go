package emqx

import (
	"github.com/fisschl/fiber/utils"
	"github.com/gofiber/fiber/v2"
	"os"
	"strings"
)

type authBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func HandleAuth(ctx *fiber.Ctx) error {
	var body authBody
	err := ctx.BodyParser(&body)
	if err != nil {
		return err
	}
	if body.Username == "default" && body.Password == os.Getenv("DEFAULT_PASSWORD") {
		return ctx.JSON(fiber.Map{
			"result": "allow",
		})
	}
	if body.Username == "public" && body.Password == "public" {
		return ctx.JSON(fiber.Map{
			"result": "allow",
		})
	}
	exists, err := utils.Rdb.Exists(ctx.Context(), body.Username).Result()
	if err == nil && exists > 0 {
		return ctx.JSON(fiber.Map{
			"result": "allow",
		})
	}
	return ctx.JSON(fiber.Map{
		"result": "deny",
	})
}

type authzBody struct {
	Username string `json:"username"`
	Topic    string `json:"topic"`
	Action   string `json:"action"`
}

func HandleAuthz(ctx *fiber.Ctx) error {
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
	if body.Username == "default" {
		return ctx.JSON(fiber.Map{
			"result": "allow",
		})
	}
	if strings.HasPrefix(body.Topic, "public") && !strings.Contains(body.Topic, "#") && !strings.Contains(body.Topic, "+") {
		return ctx.JSON(fiber.Map{
			"result": "allow",
		})
	}
	isMember, err := utils.Rdb.SIsMember(ctx.Context(), body.Username, body.Topic).Result()
	if err == nil && isMember {
		return ctx.JSON(fiber.Map{
			"result": "allow",
		})
	}
	return ctx.JSON(fiber.Map{
		"result": "deny",
	})
}
