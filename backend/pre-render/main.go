package main

import (
	"log"
	"os"

	posts "github.com/harrisbisset/blog/internal/posts"

	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
)

const root = "./blog_posts/"

func main() {
	err := os.Mkdir(root+"rendered", os.ModePerm)
	if err != nil {
		panic(err)
	}

	for _, post := range GetPostList() {
		f, err := os.Create(root + "rendered/" + post.Title)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		f.WriteString(posts.GetContentString(post))
	}
}

func GetPostList() []posts.Post {
	markdown := goldmark.New(goldmark.WithExtensions(meta.Meta))

	files, err := os.ReadDir(root + "published/")
	if err != nil {
		panic(err)
	}

	posts, errs := posts.GetFilesContents(root+"published/", files, markdown)
	if len(errs) != 0 {
		for i := range errs {
			log.Print(errs[i])
		}
	}

	return posts
}
