package main

import (
	"github.com/fisschl/fiber/emqx"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	app.Group("/emqx").
		Post("/auth", emqx.HandleAuth).
		Post("/authz", emqx.HandleAuthz)

	log.Fatal(app.Listen(":648"))
}
