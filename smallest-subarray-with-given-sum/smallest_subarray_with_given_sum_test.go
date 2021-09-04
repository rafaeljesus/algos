package smallest_subarray_with_given_sum

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	var input = []struct {
		in     []int
		s, out int
	}{
		{[]int{2, 1, 5, 2, 3, 2}, 7, 2},
		{[]int{2, 1, 5, 2, 8}, 7, 1},
		{[]int{3, 4, 1, 1, 6}, 8, 3},
	}
	for _, tt := range input {
		res := smallestSubarrayWithGivenSum(tt.in, tt.s)
		if !reflect.DeepEqual(res, tt.out) {
			t.Errorf("expected: %v,  got: %v", tt.out, res)
		}
	}
}
