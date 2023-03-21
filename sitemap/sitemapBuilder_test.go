package sitemap

import (
	"io"
	"os"
	"testing"
)

// TestBuildSitemapXML tests the BuildSitemapXML function
func TestBuildSitemapXML(t *testing.T) {
	for k, v := range testTable.buildSitemapXML {
		f, err := os.Open(k)
		if err != nil {
			t.Fatal(err)
		}
		xml, err := io.ReadAll(f)
		if err != nil {
			t.Fatal(err)
		}

		sitemapXml, err := BuildSitemapXML(v)
		if err != nil {
			t.Fatal(err)
		}

		expectedVal := string(xml)
		if expectedVal != sitemapXml {
			t.Fatalf("expected:\n%s \nbut got:\n%s", expectedVal, sitemapXml)
		}
		f.Close()
	}
}
