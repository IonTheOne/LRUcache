package lru

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	// Create a new LRU cache with capacity 3
	cache, err := NewLRUCache(3)
	if err != nil {
		t.Errorf("Failed to create LRU cache: %v", err)
	}

	// Insert key-value pairs into the cache
	cache.Put("key1", "value1")
	cache.Put("key2", "value2")
	cache.Put("key3", "value3")

	// Retrieve values from the cache
	value1, err := cache.Get("key1")
	if err != nil {
		t.Errorf("Failed to get value for key1: %v", err)
	}
	if value1 != "value1" {
		t.Errorf("Incorrect value for key1: expected value1, got %v", value1)
	}

	value2, err := cache.Get("key2")
	if err != nil {
		t.Errorf("Failed to get value for key2: %v", err)
	}
	if value2 != "value2" {
		t.Errorf("Incorrect value for key2: expected value2, got %v", value2)
	}

	value3, err := cache.Get("key3")
	if err != nil {
		t.Errorf("Failed to get value for key3: %v", err)
	}
	if value3 != "value3" {
		t.Errorf("Incorrect value for key3: expected value3, got %v", value3)
	}

	// Insert a new key-value pair, which should trigger eviction of the least recently used item
	cache.Put("key4", "value4")

	// Verify that the evicted item is no longer in the cache
	_, err = cache.Get("key1")
	if err == nil {
		t.Errorf("Expected error for key1, but got nil")
	}

	// Verify that the other items are still in the cache
	value2, err = cache.Get("key2")
	if err != nil {
		t.Errorf("Failed to get value for key2: %v", err)
	}
	if value2 != "value2" {
		t.Errorf("Incorrect value for key2: expected value2, got %v", value2)
	}

	value3, err = cache.Get("key3")
	if err != nil {
		t.Errorf("Failed to get value for key3: %v", err)
	}
	if value3 != "value3" {
		t.Errorf("Incorrect value for key3: expected value3, got %v", value3)
	}

	value4, err := cache.Get("key4")
	if err != nil {
		t.Errorf("Failed to get value for key4: %v", err)
	}
	if value4 != "value4" {
		t.Errorf("Incorrect value for key4: expected value4, got %v", value4)
	}
}
