package models

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/a-h/templ"
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
)

type Post struct {
	Name    string
	Title   string
	Order   int
	Summary string
	Date    string
	Ref     string
}

func RenderPosts() {

}

func GetPostList() []Post {
	markdown := goldmark.New(
		goldmark.WithExtensions(
			meta.Meta,
		),
	)

	files, err := os.ReadDir("blog_posts/published/")
	if err != nil {
		panic(err)
	}

	unorderedPosts, errs := GetFilesContents("blog_posts/published/", files, markdown)
	if len(errs) != 0 {
		for i := range errs {
			log.Print(errs[i])
		}
	}

	return OrderPosts(unorderedPosts)
}

func OrderPosts(posts []Post) []Post {
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

func GetFilesContents(path string, files []fs.DirEntry, markdown goldmark.Markdown) ([]Post, []error) {
	var posts []Post
	var truthyErr []error

	for _, file := range files {

		if strings.Compare(filepath.Ext(file.Name()), ".md") != 0 {
			truthyErr = append(truthyErr, fmt.Errorf("%s not markdown, %s", file.Name(), filepath.Ext(file.Name())))
			continue
		}

		data, err := os.ReadFile(fmt.Sprintf("%s%s", path, file.Name()))
		if err != nil {
			truthyErr = append(truthyErr, err)
			continue
		}

		var buf bytes.Buffer
		context := parser.NewContext()
		if err := markdown.Convert([]byte(data), &buf, parser.WithContext(context)); err != nil {
			truthyErr = append(truthyErr, err)
			continue
		}

		metaData := meta.Get(context)
		refName := strings.Replace(file.Name(), ".md", "", 1)

		posts = append(posts, Post{
			Name:    metaData["title"].(string),
			Title:   file.Name(),
			Order:   metaData["order"].(int),
			Summary: metaData["summary"].(string),
			Date:    metaData["publishedAt"].(string),
			Ref:     fmt.Sprintf("/blog_posts/published/%s", refName)})
	}

	return posts, truthyErr
}

func unsafe(html string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		_, err = io.WriteString(w, html)
		return
	})
}

func GetContent(post Post) templ.Component {
	data, err := os.ReadFile("." + post.Ref + ".md")
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}

	// removes metadata from content
	// done by second ---
	c := 0
	index := 0
	for i, j := range data[3:] {
		if j == 45 {
			c++
			if c == 3 {
				index = i
			}
		} else {
			c = 0
		}
	}
	data = data[index:]

	var buf bytes.Buffer
	if err := goldmark.Convert([]byte(data), &buf); err != nil {
		log.Fatalf("failed to convert markdown to HTML: %v", err)
	}

	return unsafe(buf.String())
}
