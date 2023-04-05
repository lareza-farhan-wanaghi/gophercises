package sitemap

import "encoding/xml"

type UrlTag struct {
	Location string `xml:"loc"`
}
type Sitemap struct {
	XMLName xml.Name `xml:"urlset"`
	Xmls    string   `xml:"xmls,attr"`
	UrlTags []UrlTag `xml:"url"`
}

// BuildSitemapXML returns an indented xml string from the urls
func BuildSitemapXML(urls []string) (string, error) {
	urlTags := []UrlTag{}
	for _, url := range urls {
		urlTags = append(urlTags, UrlTag{Location: url})
	}

	sitemap := Sitemap{
		Xmls:    "http://www.sitemaps.org/schemas/sitemap/0.9",
		UrlTags: urlTags,
	}

	result, err := xml.MarshalIndent(sitemap, "", "  ")
	if err != nil {
		return "", err
	}
	return xml.Header + string(result), nil
}
