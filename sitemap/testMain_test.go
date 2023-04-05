package sitemap

import (
	"os"
	"testing"
)

type TestTable struct {
	getDomain             map[string]string
	appendDefaultProtocol map[string]string
	crawler               map[string]string
	buildSitemapXML       map[string][]string
}

var testTable TestTable

// populateGetHostname populates test cases for the GetHostname function
func (tt *TestTable) populateGetHostname() {
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

	tt.getDomain = testCases
}

// populateAppendDefaultProtocol populates test cases for the appendDefaultProtocol function
func (tt *TestTable) populateAppendDefaultProtocol() {
	testCases := make(map[string]string)
	testCases["http://a.com/"] = "http://a.com/"
	testCases["https://a.com"] = "https://a.com"
	testCases["a.com"] = "https://a.com"
	testCases["/a"] = "https://a.com/a"
	testCases["//a.com"] = "https://a.com"
	testCases["//a.com/b/c/"] = "https://a.com/b/c/"

	tt.appendDefaultProtocol = testCases
}

// populateAppendDefaultProtocol populates test cases for the appendDefaultProtocol function
func (tt *TestTable) populateCrawler() {
	testCases := make(map[string]string)
	testCases["http://example.com/,3"] = "1"
	testCases["https://example.com/,3"] = "1"
	testCases["example.com,3"] = "1"
	testCases["https://en.wikipedia.org/wiki/Main_Page,3"] = "9"
	testCases["https://en.wikipedia.org/wiki/Main_Page,10"] = "46"

	tt.crawler = testCases
}

// populateBuildSitemapXML populates test cases for the buildSItemapXML function
func (tt *TestTable) populateBuildSitemapXML() {
	testCases := make(map[string][]string)
	testCases["ex1.xml"] = []string{
		"https://example.com",
	}
	testCases["ex2.xml"] = []string{
		"https://en.wikipedia.org/wiki/Main_Page",
		"https://en.wikipedia.org/w/index.php?title=Wikipedia:CCBYSAredirect=no",
		"https://en.wikipedia.org/w/index.php?title=Wikipedia:Contact_usoldid=967537943",
		"https://en.wikipedia.org/w/index.php?title=Wikipedia:CC_BY-SAredirect=no",
		"https://en.wikipedia.org/w/index.php?title=Wikipedia:Text_of_the_Creative_Commons_Attribution-ShareAlike_3.0_Unported_Licenseoldid=1130883578",
		"https://en.wikipedia.org/w/index.php?title=Template:Wikipedia_copyrightaction=edit",
		"https://en.wikipedia.org/wiki/Wikipedia:Contact_us",
		"https://en.wikipedia.org/w/index.php?title=Main_Pageoldid=1114291180",
		"https://en.wikipedia.org/wiki/Wikipedia:Text_of_the_Creative_Commons_Attribution-ShareAlike_3.0_Unported_License",
	}

	tt.buildSitemapXML = testCases
}

// newTestTable creates a new TestTable
func newTestTable() TestTable {
	tt := TestTable{}
	tt.populateGetHostname()
	tt.populateAppendDefaultProtocol()
	tt.populateCrawler()
	tt.populateBuildSitemapXML()
	return tt
}

// TestMain provides the entry point of the test
func TestMain(m *testing.M) {
	testTable = newTestTable()
	exitVal := m.Run()
	os.Exit(exitVal)
}
