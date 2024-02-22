package shared_test

import (
	"log"
	"os"
	"testing"

	"github.com/blog/model/shared"
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
)

func TestOrderPosts(t *testing.T) {
	posts := []shared.Post{
		{Order: 1},
		{Order: 5},
		{Order: 3},
		{Order: 2},
		{Order: 4},
	}

	expectedPosts := []shared.Post{
		{Order: 5},
		{Order: 4},
		{Order: 3},
		{Order: 2},
		{Order: 1},
	}

	orderedPosts := shared.OrderPosts(posts)

	for i := range expectedPosts {
		if expectedPosts[i].Order != orderedPosts[i].Order {
			log.Fatalf(`%dth element of orderedPosts (%d) not equal to %d`, i, orderedPosts[i].Order, expectedPosts[i].Order)
		}
	}
}

func TestGetFilesContents(t *testing.T) {
	files, err := os.ReadDir("tests/")
	if err != nil {
		log.Fatal(err)
	}
	_, errs := shared.GetFilesContents("tests/", files, goldmark.New(goldmark.WithExtensions(meta.Meta)))

	if len(errs) != 0 {
		for i := range errs {
			log.Print(errs[i])
		}
	}
	log.Fatal()
}
