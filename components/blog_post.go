package components

import (
	"github.com/gofiber/fiber/v2"
	"github.com/harrisbisset/blog/models"
	"github.com/harrisbisset/blog/render/views/view_blog"
	"github.com/harrisbisset/fnet"
	"github.com/harrisbisset/fnet/generic"
)

var BlogPost = fnet.NewComponent("Blog Post").
	Error(0, generic.GenericBuildError).
	Build()

type BlogPostWrapper struct {
	models.Post
}

func (b BlogPostWrapper) Handle(ctx *fiber.Ctx) error {

	// BlogPost.SetView(models.GetContent(b.Post))
	BlogPost.SetView(view_blog.BlogPost(b.Post))

	return BlogPost.RenderView(ctx)
}
