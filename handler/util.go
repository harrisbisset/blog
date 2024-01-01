package handler

import (
	"github.com/a-h/templ"
	"github.com/blog/view/request"
	"github.com/labstack/echo"
)

type UtilHandler struct{}

func (b UtilHandler) Handle404Show(c echo.Context) error {
	return render(c, request.NotFoundComponent())
}

func render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}
