package main

import (
	"fmt"
	"os"

	"github.com/lareza-farhan-wanaghi/gophercises/link"
)

// main provides the entry point of the app
func main() {
	reader, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer reader.Close()

	aTagFinder, err := link.NewATagFinder(reader)
	if err != nil {
		panic(err)
	}

	aTags := aTagFinder.GetATags()

	for i, aTag := range aTags {
		fmt.Printf("%d. href: '%s', text: '%s'\n", i+1, aTag.Href, aTag.Text)
	}
}
