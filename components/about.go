package components

import (
	"github.com/gofiber/fiber/v2"
	"github.com/harrisbisset/blog/render/views/view_about"
	"github.com/harrisbisset/fnet"
	"github.com/harrisbisset/fnet/generic"
)

var AboutPage = fnet.NewComponent("About Page").
	View(view_about.Show()).
	Error(0, generic.GenericBuildError).
	Build()

func AboutPageShow(ctx *fiber.Ctx) error {
	return AboutPage.RenderView(ctx)
}
