package sitemap

import (
	"strconv"
	"strings"
	"testing"
)

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

// TestGetSameDomainUrls tests the GetSameDomainUrls function
func TestGetSameDomainUrls(t *testing.T) {
	for k, v := range testTable.crawler {
		split := strings.Split(k, ",")
		depth, err := strconv.Atoi(split[1])
		if err != nil {
			t.Fatal(err)
		}

		urls, err := GetSameDomainUrls(split[0], depth)
		if err != nil {
			t.Fatal(err)
		}

		urlMapLen := strconv.Itoa(len(urls))
		if urlMapLen != v {
			t.Fatalf("wanted %s but returned %s. Url being inspected %s", v, urlMapLen, k)
		}
	}
}
