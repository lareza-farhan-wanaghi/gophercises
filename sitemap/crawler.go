package sitemap

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"sync"

	"github.com/lareza-farhan-wanaghi/gophercises/link"
)

type crawler struct {
	visitedUrl map[string]struct{}
	maxDepth   int
	rootUrl    string
	hostname   string
	wg         sync.WaitGroup
	mutex      sync.RWMutex
}

// CrawlWeb starts the crawling activity from the rootUrl to return all visited URLs for the given max depth
func (c *crawler) CrawlWeb() ([]string, error) {
	if len(c.visitedUrl) > 0 {
		return c.getUrls(), nil
	}

	c.wg.Add(1)
	go c.crawlWeb(c.rootUrl, 1)

	isDoneChan := make(chan bool)
	go func() {
		c.wg.Wait()
		isDoneChan <- true
	}()
	<-isDoneChan
	return c.getUrls(), nil
}

// crawlWeb crawls urls that concretely will visit every reachable same-domain url
func (c *crawler) crawlWeb(url string, depth int) {
	if depth > c.maxDepth {
		c.wg.Done()
		return
	}

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	c.mutex.Lock()
	c.visitedUrl[url] = struct{}{}
	c.mutex.Unlock()

	if resp.StatusCode == http.StatusOK {
		aTagFinder, err := link.NewATagFinder(resp.Body)
		if err != nil {
			panic(err)
		}
		FoundedUrls := aTagFinder.GetUrls()

		for _, FoundedUrl := range FoundedUrls {
			if linkHostname := getHostname(FoundedUrl); linkHostname == c.hostname {
				linkWithProtocol := appendDefaultProtocol(FoundedUrl, &c.hostname)
				c.mutex.RLock()
				_, ok := c.visitedUrl[linkWithProtocol]
				c.mutex.RUnlock()
				if !ok {
					c.wg.Add(1)
					go c.crawlWeb(linkWithProtocol, depth+1)
				}

			}
		}
	}

	c.wg.Done()
}

// getUrls returns a slice of the visited urls
func (c *crawler) getUrls() []string {
	result := []string{}
	for k := range c.visitedUrl {
		result = append(result, k)
	}
	return result
}

// appendDefaultProtocol return the string with http/https protocol appended to it
func appendDefaultProtocol(text string, hostname *string) string {
	protocolRegex := regexp.MustCompile(`(https://)|(http://)`)
	if !protocolRegex.MatchString(text) {
		firstTwo := text[:2]
		if firstTwo == "//" {
			text = fmt.Sprintf("%s%s", "https:", text)
		} else if firstTwo[0] == byte('/') {
			text = fmt.Sprintf("%s%s%s", "https://", *hostname, text)
		} else {
			text = fmt.Sprintf("%s%s", "https://", text)
		}
	}
	return text
}

// getHostname returns the hostname of the string
func getHostname(text string) string {
	regex := regexp.MustCompile(`(([^:\/\n?]+)\.)*([^:\/\n?]+)\.([^:\/\n?]+)`)
	result := regex.FindString(text)
	return strings.TrimSpace(result)
}

// GetSameDomainUrls visits every URL with the same domain that is found within the HTML webpage and returns the visited URLs
func GetSameDomainUrls(url string, maxDepth int) ([]string, error) {
	crawler := &crawler{
		visitedUrl: make(map[string]struct{}),
		maxDepth:   maxDepth,
		rootUrl:    appendDefaultProtocol(url, &url),
		hostname:   getHostname(url),
	}
	return crawler.CrawlWeb()
}
