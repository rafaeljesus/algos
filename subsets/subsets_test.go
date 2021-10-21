package subsets

import (
	"reflect"
	"testing"
)

func TestFindSubsets(t *testing.T) {
	var input = []struct {
		in  []int
		out [][]int
	}{
		{
			[]int{1, 3},
			[][]int{
				[]int{},
				[]int{1},
				[]int{3},
				[]int{1, 3},
			},
		},
		{
			[]int{1, 5, 3},
			[][]int{
				[]int{},
				[]int{1},
				[]int{5},
				[]int{1, 5},
				[]int{3},
				[]int{1, 3},
			},
		},
	}
	for _, tt := range input {
		res := FindSubsets(tt.in)
		if !reflect.DeepEqual(res, tt.out) {
			t.Errorf("expected: %v, got: %v", tt.out, res)
		}
	}
}
