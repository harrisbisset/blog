package handler

import (
	"github.com/blog/view/about"
	"github.com/labstack/echo"
)

type AboutHandler struct{}

func (a AboutHandler) HandleAboutShow(c echo.Context) error {
	return render(c, about.AboutBase())
}
