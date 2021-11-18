package lru_cache

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	maxSize := 3
	lru := NewLRUCache(maxSize)
	lru.Set("b", 2)
	lru.Set("a", 1)
	lru.Set("c", 3)
	if _, ok := lru.Get("a"); !ok {
		t.Fatal("expected to get 'a' key from cache")
	}
	mostRecent := lru.MostRecentlyUsed()
	if "a" != mostRecent.Key {
		t.Fatalf("expected 'a' key to be the most recent used, got: '%s'", mostRecent.Key)
	}
	lru.Set("b", 20)
	mostRecent = lru.MostRecentlyUsed()
	if "b" != mostRecent.Key {
		t.Fatalf("expected 'b' key to be the most recent used, got: '%s'", mostRecent.Key)
	}
	lru.Set("d", 4)
	expected := "c"
	if _, ok := lru.Get(expected); ok {
		t.Fatalf("expected '%s' key to be evicted", expected)
	}
}
