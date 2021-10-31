package sorting_with_heap

import (
	"testing"
)

func TestSortingBooksWithHeap(t *testing.T) {
	k := 3
	books := []*Book{
		{"Statistics without Tears: An Introduction for Non-Mathematicians", 12.68, 8.9},
		{"Operating Systems: Three Easy Pieces", 24.76, 8.9},
		{"Courage Is Calling: Fortune Favours the Brave", 13.16, 7.9},
		{"A Mind for Numbers: How to Excel at Math and Science (Even If You Flunked Algebra)", 12.19, 8.9},
		{"Probability For Dummies", 14.99, 8.1},
		{"Statistics Essentials For Dummies", 11.20, 8.1},
	}
	expected := []*Book{
		{Name: "A Mind for Numbers: How to Excel at Math and Science (Even If You Flunked Algebra)", Price: 12.19},
		{Name: "Statistics without Tears: An Introduction for Non-Mathematicians", Price: 12.68},
		{Name: "Operating Systems: Three Easy Pieces", Price: 24.76},
	}
	cheapestBooks := TopkCheapestBooks(books, k)
	for i := 0; i < len(cheapestBooks); i++ {
		if cheapestBooks[i].Price != expected[i].Price {
			t.Errorf("expected: %v, got: %v", expected[i], cheapestBooks[i])
		}
	}
}
