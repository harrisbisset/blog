package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/harrisbisset/blog/frontend/components"
	"github.com/harrisbisset/blog/frontend/models"
)

func main() {
	app := fiber.New()

	app.Get("/", components.IndexHandler{}.Handle)
	app.Get("/blog", components.BlogHandler{}.Handle)
	app.Get("/about", components.AboutHandler{}.Handle)

	for _, post := range models.GetPostList() {
		app.Get(fmt.Sprintf("/blog/%s", post.Title), components.BlogPostHandler{Post: post}.Handle)
	}

	app.Static("/dist", "./render/dist")
	app.Static("/images", "./render/images")

	app.Listen(":80")
}
