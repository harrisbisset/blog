package internal

import (
	"bytes"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

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

func GetFilesContents(path string, files []fs.DirEntry, markdown goldmark.Markdown) ([]Post, []error) {
	var blogPosts []Post
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

		blogPosts = append(blogPosts, Post{
			Name:    metaData["title"].(string),
			Title:   file.Name(),
			Order:   metaData["order"].(int),
			Summary: metaData["summary"].(string),
			Date:    metaData["publishedAt"].(string),
			Ref:     fmt.Sprintf("%s%s", path, refName)})
	}

	return blogPosts, truthyErr
}

func GetContentString(post Post) string {
	data, err := os.ReadFile(post.Ref + ".md")
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
	data = data[index+4:]

	var buf bytes.Buffer
	if err := goldmark.Convert([]byte(data), &buf); err != nil {
		log.Fatalf("failed to convert markdown to HTML: %v", err)
	}

	return buf.String()
}
