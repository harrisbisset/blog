package components

import (
	"github.com/gofiber/fiber/v2"
	"github.com/harrisbisset/blog/models"
	"github.com/harrisbisset/blog/render/views/view_index"
	"github.com/harrisbisset/fnet"
	"github.com/harrisbisset/fnet/generic"
)

var IndexPage = fnet.NewComponent("Index Page").
	View(view_index.Show(models.GetPostList())).
	Error(0, generic.GenericBuildError).
	Build()

func IndexPageShow(ctx *fiber.Ctx) error {
	return IndexPage.RenderView(ctx)
}
