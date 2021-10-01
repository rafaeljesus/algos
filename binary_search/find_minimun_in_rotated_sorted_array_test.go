package binary_search

import (
	"reflect"
	"testing"
)

func TestFindMinumunInRotatedSortedArray(t *testing.T) {
	var input = []struct {
		in  []int
		out int
	}{
		{[]int{30, 40, 50, 10, 20}, 3},
	}
	for _, tt := range input {
		res := FindMinimunInRotatedSortedArray(tt.in)
		if !reflect.DeepEqual(res, tt.out) {
			t.Errorf("expected: %v,  got: %v", tt.out, res)
		}
	}
}
