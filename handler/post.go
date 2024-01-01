package handler

import (
	. "github.com/blog/model/shared"
	"github.com/blog/view/post"
	"github.com/labstack/echo"
)

type PostHandler struct {
	Post Post
}

func (p PostHandler) HandlePostShow(c echo.Context) error {
	return render(c, post.PostBase(p.Post))
}
