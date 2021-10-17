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
	}

	for _, tt := range input {
		found, res := CheckPairWithSum(tt.in, tt.sum)
		if tt.found != found && !reflect.DeepEqual(res, tt.out) {
			t.Errorf("unxpected CheckPairWithSum result: %v, %v", found, res)
		}
	}
}
