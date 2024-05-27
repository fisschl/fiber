package main

import (
	"github.com/fisschl/fiber/amap"
	"github.com/fisschl/fiber/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func main() {
	err := utils.LoadENV()
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	app.Use(cors.New())

	amap.Register(app)

	log.Fatal(app.Listen(":8080"))
}
