package crawler

import (
	"os"
	"strconv"
	"strings"
	"testing"
)

type TestTable struct {
	getDomain             map[string]string
	appendDefaultProtocol map[string]string
	crawler               map[string]string
}

var testTable TestTable

// TestMain provides the entry point of the test
func TestMain(m *testing.M) {
	testTable = TestTable{}
	populateGetHostname()
	populateAppendDefaultProtocol()
	populateCrawler()

	exitVal := m.Run()
	os.Exit(exitVal)
}

// populateGetHostname populates test cases for the GetHostname function
func populateGetHostname() {
	testCases := make(map[string]string)
	testCases["http://a.com/"] = "a.com"
	testCases["https://a.com/"] = "a.com"
	testCases["http://a.com"] = "a.com"
	testCases["http://a.com/h"] = "a.com"
	testCases["http://a.com.c/h"] = "a.com.c"
	testCases["a.com"] = "a.com"
	testCases["a.com/h"] = "a.com"
	testCases[" a.com/h "] = "a.com"
	testCases["a.com/h.a"] = "a.com"
	testCases["https://a.com/h.a"] = "a.com"
	testCases["https://sub.a.com/h.a"] = "sub.a.com"

	testTable.getDomain = testCases
}

// populateAppendDefaultProtocol populates test cases for the appendDefaultProtocol function
func populateAppendDefaultProtocol() {
	testCases := make(map[string]string)
	testCases["http://a.com/"] = "http://a.com/"
	testCases["https://a.com"] = "https://a.com"
	testCases["a.com"] = "https://a.com"
	testCases["/a"] = "https://a.com/a"
	testCases["//a.com"] = "https://a.com"
	testCases["//a.com/b/c/"] = "https://a.com/b/c/"

	testTable.appendDefaultProtocol = testCases
}

// populateAppendDefaultProtocol populates test cases for the appendDefaultProtocol function
func populateCrawler() {
	testCases := make(map[string]string)
	testCases["http://example.com/,3"] = "1"
	testCases["https://example.com/,3"] = "1"
	testCases["example.com,3"] = "1"
	testCases["https://en.wikipedia.org/wiki/Main_Page,3"] = "9"
	testCases["https://en.wikipedia.org/wiki/Main_Page,10"] = "46"

	testTable.crawler = testCases
}

// TestGetHostname tests the getHostname function
func TestGetHostname(t *testing.T) {
	for k, v := range testTable.getDomain {
		hostname := getHostname(k)
		if hostname != v {
			t.Fatalf("Expected '%s', but returned '%s'. String on inspect: '%s'",
				v, hostname, k)
		}
	}
}

// TestAppendDefaultProtocol tests the appendDefaultProtocol function
func TestAppendDefaultProtocol(t *testing.T) {
	defaultHostname := "a.com"
	for k, v := range testTable.appendDefaultProtocol {
		appendedUrl := appendDefaultProtocol(k, &defaultHostname)
		if appendedUrl != v {
			t.Fatalf("Expected '%s', but returned '%s'. String on inspect: '%s'",
				v, appendedUrl, k)
		}
	}
}

// TestCrawler tests the crawler struct
func TestCrawler(t *testing.T) {
	for k, v := range testTable.crawler {
		split := strings.Split(k, ",")
		depth, err := strconv.Atoi(split[1])
		if err != nil {
			t.Fatal(err)
		}
		crawler := NewCrawler(split[0], depth)
		urlMap, err := crawler.CrawlWeb()
		if err != nil {
			t.Fatal(err)
		}

		urlMapLen := strconv.Itoa(len(urlMap))
		if urlMapLen != v {
			t.Fatalf("wanted %s but returned %s", v, urlMapLen)
		}
	}
}
