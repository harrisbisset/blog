package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"

	. "blog/shared"
	tem "blog/tem"

	"github.com/a-h/templ"
)

func getPages() Pages {
	home := Page{
		Name: "index",
		Ref:  "/",
		Dir:  "./index"}
	about := Page{
		Name: "about",
		Ref:  "/about",
		Dir:  "./about"}

	return Pages{home, about}
}

func main() {

	addr := os.Getenv("ADDR")

	//if it's blank, default to ":80", which means
	//listen port 80 for requests addressed to any host
	if len(addr) == 0 {
		addr = "0.0.0.0:80"
	}

	mux := http.NewServeMux()
	rh := http.RedirectHandler("http://example.org", 307)
	mux.Handle("/foo", rh)
	// mux.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
	// 	writer.WriteHeader(200)
	// 	fmt.Fprint(writer, "HELLO")
	// 	log.Printf("HELLO")
	// })

	pages := getPages()
	posts := GetPost()

	startUpProcesses(pages[0], pages) // creates static page

	for _, post := range posts {
		mux.Handle(post.Ref, templ.Handler(tem.GetBlog(post, pages)))
	}

	mux.Handle("/404", templ.Handler(tem.NotFoundComponent(), templ.WithStatus(http.StatusNotFound)))

	//create handles for pages
	// for i, page := range getPages() {
	// 	if i == 0 {
	// 		mux.Handle(page.Ref, http.FileServer(http.Dir(page.Dir)))
	// 	} else {
	// 		mux.Handle(page.Ref, templ.Handler(tem.MainComponent(page, pages)))
	// 	}
	// }

	log.Printf("server is listening at %s...", addr)
	err := http.ListenAndServe(addr, wrap(mux))
	fmt.Println(err)
}

func wrap(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// before
		h.ServeHTTP(w, r)
		// after
	})
}

func startUpProcesses(page Page, pages Pages) {
	if _, err := os.Stat(page.Dir); os.IsNotExist(err) {
		if err := os.Mkdir(page.Dir, 0755); err != nil {
			log.Fatalf("failed to create output directory: %v", err)
		}
	}
	fname := fmt.Sprintf("%s.html", page.Name)
	fmt.Println(page.Name)

	name := path.Join(page.Dir, fname)
	f, err := os.Create(name)
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}

	err = tem.CreatePage(page, pages).Render(context.Background(), f)
	if err != nil {
		log.Fatalf("failed to write page: %s", err)
	}
}
