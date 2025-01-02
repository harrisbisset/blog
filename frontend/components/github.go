package components

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/harrisbisset/blog/frontend/models"
	"github.com/harrisbisset/blog/frontend/render/views/view_index"
	render "github.com/harrisbisset/blog/internal/render"
)

type (
	LastCommitHandler struct {
		GH_TOKEN string
	}
	LastestProjectsHandler struct {
		GH_TOKEN string
	}
)

func (handler LastCommitHandler) Handle(ctx *fiber.Ctx) error {

	jsonData, err := makeGHRequest("https://api.github.com/users/harrisbisset/events/public?per_page=3", handler.GH_TOKEN)
	if err != nil {
		log.Print(err)
		return err
	}

	var publicUserEvents models.PublicUserEvent
	if err = json.Unmarshal(jsonData, &publicUserEvents); err != nil {
		log.Print(err)
		return fiber.ErrInternalServerError
	}

	return render.Render(view_index.Show(), ctx)
}

func makeGHRequest(url, GH_TOKEN string) ([]byte, error) {
	r, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fiber.ErrInternalServerError
	}

	r.Header.Add("Accept", "application/vnd.github+json")
	r.Header.Add("X-GitHub-Api-Version", "2022-11-28")
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", GH_TOKEN))

	res, err := http.DefaultClient.Do(r)
	if err != nil {
		return nil, fiber.ErrInternalServerError
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fiber.ErrInternalServerError
	}

	return resBody, nil
}
