package odlaw

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorker_NewLinkWorker(t *testing.T) {
	lw := NewLinkWorker()

	assert.IsType(t, lw, new(LinkWorker))
	assert.Equal(t, lw.container.Len(), 0)
}

func TestWorker_LinkWorkerPush(t *testing.T) {
	lw := NewLinkWorker()

	var l Link = "http://example.org"

	lw.Push(l)

	assert.Equal(t, lw.container.Len(), 1)
}

func TestWorker_LinkWorkerPushNoDuplicates(t *testing.T) {
	lw := NewLinkWorker()

	var l Link = "http://example.org"
	var l2 Link = "http://example.org"

	lw.Push(l)
	lw.Push(l2)

	assert.Equal(t, lw.container.Len(), 1)
}

func TestWorker_LinkWorkerLen(t *testing.T) {
	lw := NewLinkWorker()

	assert.Equal(t, lw.Len(), 0)

	var l Link = "http://example.org"
	var l2 Link = "http://sample.org"
	var l3 Link = "http://example.com"
	var l4 Link = "http://sample.com"

	lw.Push(l)
	lw.Push(l2)
	lw.Push(l3)
	lw.Push(l4)

	assert.Equal(t, lw.Len(), 4)
}

func TestWorker_LinkWorkerPop(t *testing.T) {
	lw := NewLinkWorker()

	assert.Equal(t, lw.Pop(), nil)

	var l Link = "http://example.org"
	lw.Push(l)

	assert.Equal(t, lw.Pop(), l)
}
