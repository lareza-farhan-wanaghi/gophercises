package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/lareza-farhan-wanaghi/gophercises/sitemap/crawler"
)

func main() {
	startTime := time.Now()

	urlFlag := flag.String("u", "https://en.wikipedia.org/wiki/Main_Page", "Specifies the url that will be inspected")
	maxDepthFlag := flag.Int("d", 10, "Specifies the max crawl depth")
	flag.Parse()

	crawler := crawler.NewCrawler(*urlFlag, *maxDepthFlag)
	urlMap, err := crawler.CrawlWeb()
	if err != nil {
		panic(err)
	}
	for url, _ := range urlMap {
		println(url)
	}
	fmt.Printf("found %d urls, completed in %v\n", len(urlMap), time.Since(startTime))

}
