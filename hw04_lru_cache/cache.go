package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (lru *lruCache) Set(key Key, value interface{}) bool {
	if _, ok := lru.Get(key); ok {
		i := lru.items[key]
		i.Value.(*cacheItem).value = value
		return true
	}

	cache := &cacheItem{key, value}
	item := lru.queue.PushFront(cache)
	lru.items[key] = item

	if lru.queue.Len() > lru.capacity {
		lru.deleteLastItem()
	}

	return false
}

func (lru *lruCache) Get(key Key) (interface{}, bool) {
	i, ok := lru.items[key]
	if !ok {
		return i, ok
	}
	lru.queue.MoveToFront(i)
	return i.Value.(*cacheItem).value, ok
}

func (lru *lruCache) Clear() {
	lru.items = make(map[Key]*ListItem, lru.capacity)
	lru.queue = NewList()
}

func (lru *lruCache) deleteLastItem() {
	lastItem := lru.queue.Back()
	lru.queue.Remove(lastItem)
	if item, ok := lastItem.Value.(*cacheItem); ok {
		delete(lru.items, item.key)
	}
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
