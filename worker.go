package odlaw

import (
	"container/list"
	"sync"
)

type Link string

type LinkWorker struct {
	container *list.List
	sync.RWMutex
}

func NewLinkWorker() *LinkWorker {
	lw := new(LinkWorker)
	lw.container = list.New()

	return lw
}

func (lw *LinkWorker) Push(l Link) {
	lw.Lock()
	defer lw.Unlock()

	for i := lw.container.Front(); i != nil; i = i.Next() {
		if i.Value == l {
			return
		}
	}

	lw.container.PushBack(l)
	return
}

func (lw *LinkWorker) Len() int {
	lw.RLock()
	defer lw.RUnlock()

	return lw.container.Len()
}

func (lw *LinkWorker) Pop() interface{} {
	lw.RLock()
	defer lw.RUnlock()

	item := lw.container.Front()
	if item != nil {
		return lw.container.Remove(item)
	}

	return nil
}
