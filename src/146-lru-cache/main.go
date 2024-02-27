package main

import "container/list"

type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
}

type pair struct {
	key, value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	if element, found := this.cache[key]; found {
		this.list.MoveToFront(element)
		return element.Value.(pair).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if element, found := this.cache[key]; found {
		this.list.MoveToFront(element)
		element.Value = pair{key, value}
	} else {
		if this.list.Len() == this.capacity {
			oldest := this.list.Back()
			delete(this.cache, oldest.Value.(pair).key)
			this.list.Remove(oldest)
		}
		this.cache[key] = this.list.PushFront(pair{key, value})
	}
}
