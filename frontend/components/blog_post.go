package components

import (
	"github.com/gofiber/fiber/v2"
	"github.com/harrisbisset/blog/frontend/render/views/view_blog"
	posts "github.com/harrisbisset/blog/internal/posts"
	render "github.com/harrisbisset/blog/internal/render"
)

type BlogPostHandler struct {
	posts.Post
}

func (b BlogPostHandler) Handle(ctx *fiber.Ctx) error {
	return render.Render(view_blog.BlogPost(b.Post), ctx)
}
