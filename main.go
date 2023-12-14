package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/a-h/templ"
)

type Post struct {
	name    string
	title   string
	summary string
	date    string
	ref     string
}

type Pages []Page

type Page struct {
	name string
	ref  string
	dir  string
}

func getPages() Pages {
	home := Page{"index", "/", "./index"}
	about := Page{"about", "/about", "./about"}

	return Pages{home, about}
}

func main() {
	pages := getPages()
	posts := getPost()

	startUpProcesses(pages[0], pages) // creates static page

	for _, post := range posts {
		http.Handle(post.ref, templ.Handler(getBlog(post, pages)))
	}

	http.Handle("/404", templ.Handler(notFoundComponent(), templ.WithStatus(http.StatusNotFound)))

	//create handles for pages
	for i, page := range getPages() {
		if i == 0 {
			http.Handle(page.ref, http.FileServer(http.Dir(page.dir)))
		} else {
			http.Handle(page.ref, templ.Handler(mainComponent(page, pages)))
		}
	}

	fmt.Println("Listening on: 8000")
	http.ListenAndServe(":8000", wrap(http.DefaultServeMux))
}

func wrap(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// before
		h.ServeHTTP(w, r)
		// after
	})
}

func startUpProcesses(page Page, pages Pages) {
	if _, err := os.Stat(page.dir); os.IsNotExist(err) {
		if err := os.Mkdir(page.dir, 0755); err != nil {
			log.Fatalf("failed to create output directory: %v", err)
		}
	}
	fname := fmt.Sprintf("%s.html", page.name)

	name := path.Join(page.dir, fname)
	f, err := os.Create(name)
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}

	err = createPage(page, pages).Render(context.Background(), f)
	if err != nil {
		log.Fatalf("failed to write page: %s", err)
	}
}
