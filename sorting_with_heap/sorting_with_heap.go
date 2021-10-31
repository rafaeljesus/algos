package sorting_with_heap

import (
	"container/heap"
)

type Book struct {
	Name          string
	Price, Rating float64
}

type Books []*Book

func (b Books) Len() int      { return len(b) }
func (b Books) Swap(i, j int) { b[i], b[j] = b[j], b[i] }

// > maxheap
// < minheap
func (b Books) Less(i, j int) bool {
	if b[i].Rating == b[j].Rating {
		return b[i].Price < b[j].Price
	}
	return b[i].Rating < b[j].Rating
}

func (b *Books) Push(x interface{}) {
	*b = append(*b, x.(*Book))
}

func (b *Books) Pop() interface{} {
	old := *b
	n := len(old)
	x := old[n-1]
	*b = old[:n-1]
	return x
}

// maxheap
// Time O(N * log(K)) | Space: O(K)
func TopkCheapestBooks(arr []*Book, k int) []*Book {
	h := &Books{}
	heap.Init(h)
	for i := 0; i < k; i++ {
		heap.Push(h, arr[i])
	}
	for i := k; i < len(arr); i++ {
		if arr[i].Rating > (*h)[0].Rating {
			// it is grater than root, take it out, and add the greater
			_ = heap.Pop(h)
			heap.Push(h, arr[i])
		}
	}
	cheapestBooks := make([]*Book, 0)
	for i := 0; i < k; i++ {
		cheapestBooks = append(cheapestBooks, heap.Pop(h).(*Book))
	}
	return cheapestBooks
}
