package main

import (
	"github.com/fisschl/fiber/dashscope"
	"github.com/fisschl/fiber/emqx"
	"github.com/fisschl/fiber/minio"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	app.Group("/oss").
		Get("/upload", minio.HandleUpload).
		Get("/download", minio.HandleDownload).
		Get("/slice", minio.HandleSlice).
		Post("/compose", minio.HandleCompose)

	app.Group("/emqx").
		Post("/auth", emqx.HandleAuth).
		Post("/authz", emqx.HandleAuthz)

	app.Group("/dashscope").
		Post("/completions", dashscope.HandleCompletions)

	log.Fatal(app.Listen(":648"))
}
