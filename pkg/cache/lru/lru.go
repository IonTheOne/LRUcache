package lru

import (
    "container/list"
    "errors"
    "sync"
)

// LRUCache is a generic LRU cache implementation
type LRUCache struct {
    capacity int
    lruList  *list.List
    items    map[interface{}]*list.Element
    mutex    sync.Mutex
}

// entry represents a cache entry
type entry struct {
    key   interface{}
    value interface{}
}

// NewLRUCache creates a new LRU cache with the given capacity
func NewLRUCache(capacity int) (*LRUCache, error) {
    if capacity <= 0 {
        return nil, errors.New("capacity must be a positive integer")
    }
    return &LRUCache{
        capacity: capacity,
        lruList:  list.New(),
        items:    make(map[interface{}]*list.Element),
    }, nil
}

// Get retrieves the value associated with the key from the cache
func (c *LRUCache) Get(key interface{}) (interface{}, error) {
    c.mutex.Lock()
    defer c.mutex.Unlock()

    if elem, ok := c.items[key]; ok {
        c.lruList.MoveToFront(elem)
        return elem.Value.(*entry).value, nil
    }
    return nil, errors.New("key not found in cache")
}

// Put inserts a key-value pair into the cache
func (c *LRUCache) Put(key, value interface{}) {
    c.mutex.Lock()
    defer c.mutex.Unlock()

    if elem, ok := c.items[key]; ok {
        c.lruList.MoveToFront(elem)
        elem.Value.(*entry).value = value
    } else {
        if len(c.items) >= c.capacity {
            c.evict()
        }
        elem := c.lruList.PushFront(&entry{key, value})
        c.items[key] = elem
    }
}

// evict removes the least recently used item from the cache
func (c *LRUCache) evict() {
    if back := c.lruList.Back(); back != nil {
        entry := c.lruList.Remove(back).(*entry)
        delete(c.items, entry.key)
    }
}
