package queue4go

import (
	"sync"
)

type QueueTable struct {
	sync.RWMutex
	items []*QueueItem
}

func (this *QueueTable) Length() int {
	this.RLock()
	defer this.RUnlock()
	return len(this.items)
}

func (this *QueueTable) Pop() interface{} {
	this.Lock()
	defer this.Unlock()
	value := this.items[0].Data()
	this.items = this.items[1:]
	return value
}

func (this *QueueTable) Push(item interface{}) bool {
	this.Lock()
	defer this.Unlock()
	data := NewQueueItem(item)
	this.items = append(this.items, data)
	return data == this.items[0]
}
