package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/harrisbisset/blog/components"
	"github.com/harrisbisset/blog/models"
)

func main() {

	models.RenderPosts()

	app := fiber.New()

	app.Get("/", components.IndexPageShow)
	app.Get("/about", components.AboutPageShow)

	for _, post := range models.GetPostList() {
		app.Get(fmt.Sprintf("/blog/%s", post.Name[:len(post.Name)-3]), components.BlogPostWrapper{Post: post}.Handle)
	}

	app.Static("/dist", "./render/dist")
	app.Static("/images", "./render/images")

	app.Listen(":80")
}
