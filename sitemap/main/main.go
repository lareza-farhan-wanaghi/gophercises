package main

import (
	"flag"
	"fmt"

	"github.com/lareza-farhan-wanaghi/gophercises/sitemap"
)

// main provides the entry point of the app
func main() {
	urlFlag := flag.String("u", "https://en.wikipedia.org/wiki/Main_Page", "Specifies the url that will be inspected")
	maxDepthFlag := flag.Int("d", 3, "Specifies the max crawl depth")
	flag.Parse()

	urls, err := sitemap.GetSameDomainUrls(*urlFlag, *maxDepthFlag)
	if err != nil {
		panic(err)
	}

	for _, url := range urls {
		fmt.Println(url)
	}
	sitemapXml, err := sitemap.BuildSitemapXML(urls)
	if err != nil {
		panic(err)
	}
	fmt.Println(sitemapXml)
}
