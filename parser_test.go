package odlaw

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

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
