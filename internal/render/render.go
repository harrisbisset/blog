package internal

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func Render(tc templ.Component, ctx *fiber.Ctx) error {
	return adaptor.HTTPHandler(templ.Handler(tc))(ctx)
}
