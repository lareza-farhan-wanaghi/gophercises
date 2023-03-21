package link

import (
	"os"
	"testing"
)

// TestATagFinderGetATags tests the GetATags function of ATagFinder struct
func TestATagFinderGetATags(t *testing.T) {
	for k, v := range testTable.aTagFinder {
		reader, err := os.Open(k)
		if err != nil {
			panic(err)
		}
		aTagFinder, err := NewATagFinder(reader)
		if err != nil {
			t.Fatal(err)
		}

		aTags := aTagFinder.GetATags()

		for i, aTag := range aTags {
			if !aTag.Similar(v[i]) {
				t.Fatalf("Expected '%s':'%s', but returned '%s':'%s'. File on test: %s",
					v[i].Href, v[i].Text, aTag.Href, aTag.Text, k)
			}
		}
		reader.Close()
	}
}

// TestATagFinderGetUrls tests the GetUrls function of ATagFinder struct
func TestATagFinderGetUrls(t *testing.T) {
	for k, v := range testTable.aTagFinder {
		reader, err := os.Open(k)
		if err != nil {
			panic(err)
		}
		aTagFinder, err := NewATagFinder(reader)
		if err != nil {
			t.Fatal(err)
		}

		urls := aTagFinder.GetUrls()

		for i, url := range urls {
			if url != v[i].Href {
				t.Fatalf("Expected '%s', but returned '%s'. File on test: %s",
					v[i].Href, url, k)
			}
		}
		reader.Close()
	}
}

// TestConvertToString tests the convertToString function
func TestConvertToString(t *testing.T) {
	for k, v := range testTable.convertToString {
		convertedString := convertToString([]byte(k))

		if convertedString != v {
			t.Fatalf("Expected '%s', but returned '%s'. String on inspect: '%s'",
				v, convertedString, k)
		}

	}
}
