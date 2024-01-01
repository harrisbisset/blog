package handler

import (
	. "github.com/blog/model/shared"
	"github.com/blog/view/blog"
	"github.com/labstack/echo"
)

type BlogHandler struct {
	Posts []Post
}

func (i BlogHandler) HandleBlogShow(c echo.Context) error {
	return render(c, blog.BlogBase(i.Posts))
}
