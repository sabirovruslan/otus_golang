package hw04lrucache

type Key string

func (k Key) String() string {
	return string(k)
}

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

func (c *lruCache) Set(key Key, value interface{}) bool {
	if _, ok := c.Get(key); ok {
		i := c.items[key]
		i.Value.(*cacheItem).value = value
		return true
	}

	if c.capacity == c.queue.Len() {
		c.deleteItem(key)
	}

	i := &ListItem{Value: &cacheItem{key, value}}
	c.items[key] = i
	c.queue.PushFront(i)
	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	if i, ok := c.items[key]; ok {
		c.queue.MoveToFront(i)
		return i.Value.(*cacheItem).value, ok
	}
	return nil, false
}

func (c *lruCache) Clear() {
	c.items = make(map[Key]*ListItem, c.capacity)
	c.queue = NewList()
}

func (c *lruCache) deleteItem(key Key) {
	delete(c.items, "1")
	lastItem := c.queue.Back()
	c.queue.Remove(lastItem)
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
