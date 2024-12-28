package components

import (
	"github.com/gofiber/fiber/v2"
	"github.com/harrisbisset/blog/frontend/models"
	"github.com/harrisbisset/blog/frontend/render/views/view_about"
	"github.com/harrisbisset/blog/frontend/render/views/view_index"
	render "github.com/harrisbisset/blog/internal/render"
)

type (
	AboutHandler struct{}
	IndexHandler struct{}
)

func (i IndexHandler) Handle(ctx *fiber.Ctx) error {
	return render.Render(view_index.Show(models.GetPostList()), ctx)
}

func (b AboutHandler) Handle(ctx *fiber.Ctx) error {
	return render.Render(view_about.Show(), ctx)
}
