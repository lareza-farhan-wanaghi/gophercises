package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/lareza-farhan-wanaghi/gophercises/sitemap"
)

// main provides the entry point of the app
func main() {
	maxDepthFlag := flag.Int("d", 2, "Specifies the max crawl depth")
	flag.Parse()

	targetRootUrl := os.Args[1]
	urls, err := sitemap.GetSameDomainUrls(targetRootUrl, *maxDepthFlag)
	if err != nil {
		panic(err)
	}

	sitemapXml, err := sitemap.BuildSitemapXML(urls)
	if err != nil {
		panic(err)
	}

	outPath := os.Args[2]
	outFile, err := os.OpenFile(outPath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	outFile.WriteString(sitemapXml)

	fmt.Printf("Created the sitemap XML for %s with depth of %d at %s\n", targetRootUrl, *maxDepthFlag, outPath)
}
