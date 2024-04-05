# LRU cache 

## Description

This implementation creates a Cache struct that holds a map of items and a doubly linked list (lruList) to maintain the order of items based on their access time. The NewCache function initializes a new cache with a given capacity. The Get method retrieves the value associated with a key from the cache, and the Put method inserts a new key-value pair into the cache. When the cache is full and a new item needs to be inserted, the least recently used item is evicted from the cache.
