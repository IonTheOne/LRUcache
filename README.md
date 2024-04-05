# LRU cache 

## Description

This implementation creates a Cache struct that holds a map of items and a doubly linked list (lruList) to maintain the order of items based on their access time. The NewCache function initializes a new cache with a given capacity. The Get method retrieves the value associated with a key from the cache, and the Put method inserts a new key-value pair into the cache. When the cache is full and a new item needs to be inserted, the least recently used item is evicted from the cache.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

What things you need to install the software and how to install them.

### Installing

A step by step series of examples that tell you how to get a development environment running.

## Running the application

To build the application, run:

```bash
make build
```

For another make comands see Makefile to run tests and generate documentation, etc.
