package twosum

import (
	"reflect"
	"testing"
)

func TestCheckPairWithSum(t *testing.T) {
	var input = []struct {
		in    []int
		sum   int
		found bool
		out   []int
	}{
		{[]int{0, -1, 2, -3, 1}, -2, true, []int{-3, 1}},
		{[]int{1, -2, 1, 0, 5}, 0, false, []int{}},
		{[]int{1, 4, 45, 6, 10, 8}, 16, true, []int{6, 10}},
	}

	for _, tt := range input {
		found, res := CheckPairWithSum1(tt.in, tt.sum)
		if tt.found != found {
			t.Errorf("unxpected CheckPairWithSum1 result: %v", found)
		}
		if !reflect.DeepEqual(res, tt.out) {
			t.Errorf("unxpected CheckPairWithSum1 result: %v", res)
		}
	}
}
