package components

import (
	"github.com/gofiber/fiber/v2"
	"github.com/harrisbisset/blog/frontend/models"
	"github.com/harrisbisset/blog/frontend/render/views/view_blog"
)

type BlogPostHandler struct {
	models.Post
}

func (b BlogPostHandler) Handle(ctx *fiber.Ctx) error {
	return render(view_blog.BlogPost(b.Post), ctx)
}
