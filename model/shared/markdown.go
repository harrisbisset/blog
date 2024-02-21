package shared

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"strings"

	"github.com/a-h/templ"
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
)

func GetPosts() ([]Post, error) {
	markdown := goldmark.New(
		goldmark.WithExtensions(
			meta.Meta,
		),
	)

	files, err := os.ReadDir("posts/MD/")
	if err != nil {
		return nil, err
	}

	unorderedPosts := getFilesContents(files, markdown)
	orderedPosts := orderPosts(unorderedPosts)

	return orderedPosts, nil
}

func orderPosts(posts []Post) []Post {
	count := 1
	for count < len(posts) {
		item := posts[count-1]
		count2 := count

		for count2 > 0 && posts[count2-2].Order > item.Order {
			posts[count2-1] = posts[count2-2]
			count2 -= 1
		}

		posts[count2-1] = item
		count += 1
	}

	return posts
}

func getFilesContents(files []fs.DirEntry, markdown goldmark.Markdown) []Post {
	var posts []Post

	for _, file := range files {

		if strings.Compare(file.Name()[len(file.Name())-2:], ".md") != 0 {
			log.Print("file not markdown")
			continue
		}

		data, err := os.ReadFile(fmt.Sprintf("posts/MD/%s", file.Name()))
		if err != nil {
			log.Print(err)
			continue
		}

		var buf bytes.Buffer
		context := parser.NewContext()
		if err := markdown.Convert([]byte(data), &buf, parser.WithContext(context)); err != nil {
			log.Print(err)
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
			Ref:     fmt.Sprintf("/posts/MD/%s", refName)})
	}

	return posts
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
