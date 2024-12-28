package models

import (
	"context"
	"io"
	"log"
	"os"

	posts "github.com/harrisbisset/blog/internal/posts"

	"github.com/a-h/templ"
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
)

func GetPostList() []posts.Post {
	markdown := goldmark.New(goldmark.WithExtensions(meta.Meta))

	files, err := os.ReadDir("./render/posts/published/")
	if err != nil {
		panic(err)
	}

	unorderedPosts, errs := posts.GetFilesContents("./render/posts/published/", files, markdown)
	if len(errs) != 0 {
		for i := range errs {
			log.Print(errs[i])
		}
	}

	return OrderPosts(unorderedPosts)
}

func OrderPosts(posts []posts.Post) []posts.Post {
	count := 1
	for count < len(posts) {
		item := posts[count]
		count2 := count

		for count2 > 0 && posts[count2-1].Order < item.Order {
			posts[count2] = posts[count2-1]
			count2 -= 1
		}

		posts[count2] = item
		count += 1
	}

	return posts
}

func GetContent(post posts.Post) templ.Component {
	data, err := os.ReadFile("./render/posts/rendered/" + post.Title)
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}

	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		_, err = io.WriteString(w, string(data))
		return
	})
}
