package shared

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/a-h/templ"
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
)

func GetPost() []Post {
	var postBs []Post
	var ord []int
	var tPost []Post

	markdown := goldmark.New(
		goldmark.WithExtensions(
			meta.Meta,
		),
	)

	files, err := os.ReadDir("blogPosts/")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		data, err := os.ReadFile(fmt.Sprintf("blogPosts/%s", file.Name()))
		if err != nil {
			log.Panicf("failed reading data from file: %s", err)
		}

		var buf bytes.Buffer
		context := parser.NewContext()
		if err := markdown.Convert([]byte(data), &buf, parser.WithContext(context)); err != nil {
			panic(err)
		}

		// need to handle if file in blogPosts that isn't markdown
		// fails on {interface}.(type)

		metaData := meta.Get(context)
		ord = append(ord, metaData["order"].(int))

		name := strings.Replace(file.Name(), ".md", "", 1)
		tPost = append(tPost, Post{
			Name:    metaData["title"].(string),
			Title:   file.Name(),
			Summary: metaData["summary"].(string),
			Date:    metaData["publishedAt"].(string),
			Ref:     fmt.Sprintf("/blogPosts/%s", name)})
	}

	// orders array
	for i := len(ord) - 1; i >= 0; i-- {
		postBs = append(postBs, tPost[i])
	}

	return postBs
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
	c := 0
	index := 0
	for i, j := range data {
		if j == 45 {
			c++
			if c == 6 {
				index = i
			}
		}
	}
	data = data[index:]

	var buf bytes.Buffer
	if err := goldmark.Convert([]byte(data), &buf); err != nil {
		log.Fatalf("failed to convert markdown to HTML: %v", err)
	}

	return unsafe(buf.String())
}
