package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/lareza-farhan-wanaghi/gophercises/link"
)

// main provides the entry point of the app
func main() {
	fileFlag := flag.String("f", "ex1.html", "Specifies the file containing the html")
	flag.Parse()

	reader, err := os.Open(*fileFlag)
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
