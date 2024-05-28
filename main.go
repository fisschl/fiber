package main

import (
	"github.com/fisschl/fiber/oss"
	"github.com/fisschl/fiber/utils"
	"log"
)

func main() {
	app := utils.NewFiberApp()
	oss.RegisterRouter(app)
	log.Fatal(app.Listen(":648"))
}
