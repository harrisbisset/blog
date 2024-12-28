package components

import (
	"github.com/gofiber/fiber/v2"
	"github.com/harrisbisset/blog/frontend/models"
	"github.com/harrisbisset/blog/frontend/render/views/view_about"
	"github.com/harrisbisset/blog/frontend/render/views/view_index"
)

type (
	AboutHandler struct{}
	IndexHandler struct{}
)

func (i IndexHandler) Handle(ctx *fiber.Ctx) error {
	return render(view_index.Show(models.GetPostList()), ctx)
}

func (b AboutHandler) Handle(ctx *fiber.Ctx) error {
	return render(view_about.Show(), ctx)
}
