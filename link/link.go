package link

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

type ATag struct {
	Href string `xml:"loc"`
	Text string `xml:"-"`
}

// similar determines whether the link is similar to the other
func (link *ATag) Similar(other *ATag) bool {
	if link.Href != other.Href {
		return false
	}
	if link.Text != other.Text {
		return false
	}
	return true
}

type ATagFinder struct {
	HtmlReader     io.Reader
	tokenizer      *html.Tokenizer
	collectedATags []*ATag
	tempATags      []*ATag
}

// parseHTML parses HTML data into a slice of aTag structs
func (h *ATagFinder) parseHTML() error {
	h.tokenizer = html.NewTokenizer(h.HtmlReader)
	h.tempATags = make([]*ATag, 0)
	h.collectedATags = make([]*ATag, 0)
	for {
		tokenType := h.tokenizer.Next()
		switch tokenType {
		case html.ErrorToken:
			h.tempATags = nil
			return nil
		case html.TextToken:
			h.handleTextToken()
		case html.StartTagToken:
			h.handleStartTagToken()
		case html.EndTagToken:
			h.handleEndTagToken()
		}
	}
}

// handleTextToken handles text tag tokens while parsing the html file
func (h *ATagFinder) handleTextToken() {
	text := convertToString(h.tokenizer.Text())
	if text != "" {
		for _, link := range h.tempATags {
			if link.Text == "" {
				link.Text = text
			} else {
				link.Text += fmt.Sprintf(" %s", text)
			}
		}
	}
}

// handleStartTagToken handles start tag tokens while parsing the html file
func (h *ATagFinder) handleStartTagToken() {
	tagName, _ := h.tokenizer.TagName()
	if convertToString(tagName) == "a" {
		href := ""
		for {
			k, v, _ := h.tokenizer.TagAttr()
			key := convertToString(k)
			if key == "" {
				break
			} else if key == "href" {
				href = convertToString(v)
				break
			}
		}
		h.tempATags = append(h.tempATags, &ATag{Href: href})
	}
}

// handleEndTagToken handles end tag tokens while parsing the html file
func (h *ATagFinder) handleEndTagToken() {
	tagName, _ := h.tokenizer.TagName()
	if convertToString(tagName) == "a" && len(h.tempATags) > 0 {
		h.collectedATags = append(h.collectedATags, h.tempATags[len(h.tempATags)-1])
		h.tempATags = h.tempATags[:len(h.tempATags)-1]
	}
}

// GetATags retrieves the ATagFinder's collected aTags
func (h *ATagFinder) GetATags() []*ATag {
	return h.collectedATags
}

// GetUrls retrieves urls from the ATagFinder's collected aTags
func (h *ATagFinder) GetUrls() []string {
	result := []string{}
	for _, aTag := range h.collectedATags {
		result = append(result, aTag.Href)
	}
	return result
}

// ConvertToString converts bytes into a string with spaces and new lines have been trimmed
func convertToString(text []byte) string {
	result := string(text)
	result = strings.TrimSpace(result)
	result = strings.TrimRight(result, "\r\n")
	return result
}

// NewATagFinder creates a ATagFinder that have been parsed the html file
func NewATagFinder(reader io.Reader) (ATagFinder, error) {
	parser := ATagFinder{HtmlReader: reader}
	err := parser.parseHTML()
	if err != nil {
		return parser, err
	}
	return parser, nil
}
