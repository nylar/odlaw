package odlaw

import (
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
)

func TestParser_NewDocument(t *testing.T) {
	html := `<html><body><p>Hello, World!</p></body></html>`

	doc := NewDocument(html)

	assert.IsType(t, new(goquery.Document), doc)
}

func TestParser_NewDocumentStripsJunk(t *testing.T) {
	html := `
<!DOCTYPE html>
<html>
<head>
	<title>Hello</title>

	<script type="text/javascript">
	alert("Hello, World");
	</script>

	<style>
	* { font-family: 'Comic Sans' }
	</style>
</head>

<body>
	<p>Hello, World!</p>
</body>
</html>`

	doc := NewDocument(html)
	js := doc.Find("script")
	css := doc.Find("style")
	p := doc.Find("p")

	// Should be removed and thus be 0 matching nodes.
	assert.Equal(t, js.Length(), 0)
	assert.Equal(t, css.Length(), 0)

	// Everything else should be left as is.
	assert.Equal(t, p.Length(), 1)
}

func TestParser_ExtractTitleFromTitle(t *testing.T) {
	html := `
<!DOCTYPE html>
<html>
<head>
	<title>The Title</title>
</head>
<body></body>
</html>`

	doc := NewDocument(html)

	title := ExtractTitle(doc)

	assert.Equal(t, "The Title", title)
}

func TestParser_ExtractTitleFromHeading1(t *testing.T) {
	html := `
<!DOCTYPE html>
<html>
<head></head>
<body>
	<h1>The Title</h1>
</body>
</html>`

	doc := NewDocument(html)

	title := ExtractTitle(doc)

	assert.Equal(t, "The Title", title)
}

func TestParse_ExtractTitlePrecendence(t *testing.T) {
	html := `
<!DOCTYPE html>
<html>
<head>
	<title>About Us</title>
</head>
<body>
	<h1>We rock</h1>
</body>
</html>`

	doc := NewDocument(html)

	title := ExtractTitle(doc)

	assert.Equal(t, "About Us", title)
	assert.NotEqual(t, "We Rock", title)
	assert.NotEqual(t, "", title)
}

func TestParser_ExtractTitleEmpty(t *testing.T) {
	html := `<!DOCTYPE html><html><head></head><body></body></html>`

	doc := NewDocument(html)

	title := ExtractTitle(doc)

	assert.Equal(t, "", title)
}

func TestParser_ExtractLinks_Empty(t *testing.T) {
	links := ExtractLinks("")

	assert.Equal(t, len(links), 0)
}

func TestParser_ExtractLinks_Valid(t *testing.T) {
	htmlSoup := `
<p>
	<a href="http://example.org/1">Link 1</a>
	<br>
	<a href="http://example.org/2">Link 2</a>
</p>`

	links := ExtractLinks(htmlSoup)

	assert.Equal(t, len(links), 2)
}

func TestParser_ExtractLinks_Invalid(t *testing.T) {
	// This should return an error but html.Parse doesn't seem to care.
	invalidHtml := `<html><body><aef<eqf>>>qq></body></ht>`
	links := ExtractLinks(invalidHtml)

	assert.Equal(t, len(links), 0)
}

func TestParser_ExtractLinks_NoDuplicates(t *testing.T) {
	htmlWithDupes := `
<p>
	<a href="http://example.org/1">Link 1</a>
	<a href="http://example.org/2">Link 1</a>
	<a href="http://example.org/3">Link 1</a>
	<a href="http://example.org/1">Link 1</a>
</p>`

	links := ExtractLinks(htmlWithDupes)

	assert.Equal(t, len(links), 3)
}
