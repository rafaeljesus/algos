package topk_numbers

import (
	"container/heap"
	"sort"
)

// https://pkg.go.dev/container/heap#example-package-IntHeap
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// Given an unsorted array of numbers, find the 'K' largest numbers in it.
// in: [3, 1, 5, 12, 2, 11], K = 3
// out: [5, 12, 11]
// Time: O(logN) | Space: O(1)
func TopkNumbers(nums []int, k int) []int {
	h := &IntHeap{}
	heap.Init(h)
	for i := 0; i < k; i++ {
		heap.Push(h, nums[i])
	}
	for i := k; i < len(nums); i++ {
		min := (*h)[0]
		if nums[i] > min {
			_ = heap.Pop(h)
			heap.Push(h, nums[i])
		}
	}
	return (*h)[:k]
}

// Time: O(N * logN) | Space: O(1)
func TopkNumbers2(nums []int, k int) []int {
	sort.Ints(nums)
	return nums[k:]
}
