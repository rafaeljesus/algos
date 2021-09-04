package lru_cache

type LRUCache struct {
	index            map[string]*DoublyLinkedListNode
	maxSize          int
	currentSize      int
	listOfMostRecent *DoublyLinkedList
}
