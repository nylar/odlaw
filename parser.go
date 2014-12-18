package odlaw

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Extracts all anchors (with href attributes) from a document and return a list
// of the anchors. Should return an error but goquery.NewDocumentFromReader that
// subsequently calls html.Parse doesn't like returning errors for bad markup.
func ExtractLinks(document string) []string {
	links := []string{}
	linkTracker := make(map[string]bool)
	
	// Skip the error as no combination of invalid HTML will trigger an error.
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(document))
	
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		// Only interested in anchors that have a href attribute.
		link, href := s.Attr("href")
		if href {
			if _, ok := linkTracker[link]; !ok {
				links = append(links, link)
				linkTracker[link] = true
			}
		}
	})
	
	return links
}
