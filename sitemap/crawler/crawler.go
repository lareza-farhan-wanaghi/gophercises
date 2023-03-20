package crawler

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"sync"

	"github.com/lareza-farhan-wanaghi/gophercises/link/parser"
)

type Crawler struct {
	visitedUrl map[string]struct{}
	maxDepth   int
	rootUrl    string
	hostname   string
	wg         sync.WaitGroup
	mutex      sync.RWMutex
}

// CrawlWeb crawls the rootUrl and its neighbors to return all visited URLs for the given max depth
func (c *Crawler) CrawlWeb() (map[string]struct{}, error) {
	c.wg.Add(1)
	go c.crawlWeb(c.rootUrl, 1)

	isDoneChan := make(chan bool)
	go func() {
		c.wg.Wait()
		isDoneChan <- true
	}()
	<-isDoneChan
	return c.visitedUrl, nil
}

// crawlWeb is the inner function of CrawlWeb that concretely will visit reachable URLs
func (c *Crawler) crawlWeb(url string, depth int) {
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
		links, err := parser.ParseHTML(resp.Body)
		if err != nil {
			panic(err)
		}

		for _, link := range links {
			if linkHostname := getHostname(link.Href); linkHostname == c.hostname {
				linkWithProtocol := appendDefaultProtocol(link.Href, &c.hostname)
				c.mutex.RLock()
				if _, ok := c.visitedUrl[linkWithProtocol]; !ok {
					c.wg.Add(1)
					go c.crawlWeb(linkWithProtocol, depth+1)
				}
				c.mutex.RUnlock()
			}
		}
	}

	c.wg.Done()
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

// NewCrawler creates a crawler object with the defined specifications
func NewCrawler(url string, maxDepth int) *Crawler {
	return &Crawler{
		visitedUrl: make(map[string]struct{}),
		maxDepth:   maxDepth,
		rootUrl:    appendDefaultProtocol(url, &url),
		hostname:   getHostname(url),
	}
}
