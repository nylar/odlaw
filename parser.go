package odlaw

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// NewDocument generates a new Document and strips out useless tags.
func NewDocument(document string) *goquery.Document {
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(document))

	doc.Find("style, script, link, iframe, frame, embed").Remove()

	return doc
}

// ExtractTitle searches for a title tag or the first h1 tag, if none of these
// are found then just set the title to an empty string.
func ExtractTitle(doc *goquery.Document) string {
	title := doc.Find("title")
	if title.Length() > 0 {
		return title.First().Text()
	}

	heading := doc.Find("h1")
	if heading.Length() > 0 {
		return heading.First().Text()
	}

	return ""
}

// ExtractAuthor searches for an author in a meta tag or an author class or ID.
func ExtractAuthor(doc *goquery.Document) string {
	meta := doc.Find("meta[name=author]")
	if meta.Length() > 0 {
		author, exists := meta.Attr("content")
		if exists {
			return author
		}
	}

	// FIXME: This will not adapt well, some sites are going to use funky
	//       and exotic class names.
	author := doc.Find(".author, #author")
	if author.Length() > 0 {
		return author.First().Text()
	}

	return ""
}

// ExtractLinks all anchors (with href attributes) from a document and return a list
// of the anchors. Should return an error but goquery.NewDocumentFromReader that
// subsequently calls html.Parse doesn't like returning errors for bad markup.
func ExtractLinks(doc *goquery.Document) []string {
	links := []string{}
	linkTracker := make(map[string]bool)

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

// ExtractText extracts all p tags from a page.
func ExtractText(doc *goquery.Document) string {
	texts := []string{}
	doc.Find("p").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		if text != "" {
			texts = append(texts, text)
		}
	})
	return strings.Join(texts, "\n")
}
