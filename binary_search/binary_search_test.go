package binary_search

import (
	"reflect"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	var input = []struct {
		in          []int
		target, out int
	}{
		{[]int{1, 3, 5, 7, 8}, 5, 2},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 0, -1},
		{[]int{2, 8, 89, 120, 1000}, 120, 3},
	}
	for _, tt := range input {
		res := BinarySearch(tt.in, tt.target)
		if !reflect.DeepEqual(res, tt.out) {
			t.Errorf("expected: %v,  got: %v", tt.out, res)
		}
		res = BinarySearchRecursive(tt.in, tt.target, 0, len(tt.in)-1)
		if !reflect.DeepEqual(res, tt.out) {
			t.Errorf("expected: %v,  got: %v", tt.out, res)
		}
	}
}
