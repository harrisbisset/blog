package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/blog/handler"
	"github.com/blog/model/shared"
	"github.com/labstack/echo"
)

func main() {
	blogPosts := shared.GetPosts()
	app := echo.New()

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = ":3000"
	}

	s := http.Server{
		Addr:    port,
		Handler: app,
	}

	app.Static("/dist", "dist")
	app.Static("/public", "public")

	app.GET("/404", handler.UtilHandler{}.Handle404Show)
	app.GET("/about", handler.AboutHandler{}.HandleAboutShow)
	app.GET("/blog", handler.BlogHandler{Posts: blogPosts}.HandleBlogShow)

	for _, post := range blogPosts {
		app.GET(fmt.Sprintf("/blog/%s", post.Name[:len(post.Name)-3]), handler.PostHandler{Post: post}.HandlePostShow)
	}

	log.Printf("server is listening at %s...", port)
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
