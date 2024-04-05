package lru

import (
    "container/list"
    "errors"
)

// Cache is a generic LRU cache implementation
type Cache struct {
    capacity int
    lruList  *list.List
    items    map[interface{}]*list.Element
}

// Entry represents a cache entry
type Entry struct {
    key   interface{}
    value interface{}
}

// NewCache creates a new LRU cache with the given capacity
func NewCache(capacity int) (*Cache, error) {
    if capacity <= 0 {
        return nil, errors.New("capacity must be a positive integer")
    }
    return &Cache{
        capacity: capacity,
        lruList:  list.New(),
        items:    make(map[interface{}]*list.Element),
    }, nil
}

// Get retrieves the value associated with the key from the cache
func (c *Cache) Get(key interface{}) (interface{}, error) {
    if elem, ok := c.items[key]; ok {
        c.lruList.MoveToFront(elem)
        return elem.Value.(*Entry).value, nil
    }
    return nil, errors.New("key not found in cache")
}

// Put inserts a key-value pair into the cache
func (c *Cache) Put(key, value interface{}) {
    if elem, ok := c.items[key]; ok {
        c.lruList.MoveToFront(elem)
        elem.Value.(*Entry).value = value
    } else {
        if len(c.items) >= c.capacity {
            c.evict()
        }
        elem := c.lruList.PushFront(&Entry{key, value})
        c.items[key] = elem
    }
}

// evict removes the least recently used item from the cache
func (c *Cache) evict() {
    if back := c.lruList.Back(); back != nil {
        entry := c.lruList.Remove(back).(*Entry)
        delete(c.items, entry.key)
    }
}
