package lru_cache

import (
	"container/list"
)

type Entry struct {
	Key   string
	Value int
}

type LRUCache struct {
	MaxSize    int
	Cache      map[string]*list.Element
	LinkedList *list.List
}

func NewLRUCache(maxSize int) *LRUCache {
	return &LRUCache{
		MaxSize:    maxSize,
		Cache:      make(map[string]*list.Element, maxSize),
		LinkedList: list.New(),
	}
}

func (lru *LRUCache) Get(key string) (int, bool) {
	e, ok := lru.Cache[key]
	if !ok {
		return -1, false
	}
	lru.LinkedList.MoveToFront(e)
	return e.Value.(*Entry).Value, true
}

func (lru *LRUCache) Set(key string, value int) {
	if e, ok := lru.Cache[key]; ok {
		lru.LinkedList.MoveToFront(e)
		e.Value.(*Entry).Value = value
		return
	}
	e := lru.LinkedList.PushFront(&Entry{Key: key, Value: value})
	lru.Cache[key] = e
	if lru.LinkedList.Len() > lru.MaxSize {
		tail := lru.LinkedList.Back()
		lru.LinkedList.Remove(tail)
		delete(lru.Cache, tail.Value.(*Entry).Key)
	}
}

func (lru *LRUCache) MostRecentlyUsed() *Entry {
	return lru.LinkedList.Front().Value.(*Entry)
}
