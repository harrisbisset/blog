package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/harrisbisset/blog/backend/client/components"
)

func main() {
	app := fiber.New()

	app.Get("/", components.IndexHandler{}.Handle)

	app.Static("/dist", "./render/dist")

	app.Listen(":8080")
}
