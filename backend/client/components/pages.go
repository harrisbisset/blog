package components

import (
	"github.com/gofiber/fiber/v2"
)

type (
	IndexHandler struct{}
)

func (i IndexHandler) Handle(ctx *fiber.Ctx) error {
	return nil
}
