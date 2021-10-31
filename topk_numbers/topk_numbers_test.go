package topk_numbers

import (
	"reflect"
	"testing"
)

func TestTopkNumbers(t *testing.T) {
	var tests = []struct {
		in  []int
		k   int
		out []int
	}{
		{[]int{3, 1, 5, 12, 2, 11}, 3, []int{5, 12, 11}},
	}
	for _, tt := range tests {
		res := TopkNumbers(tt.in, tt.k)
		if !reflect.DeepEqual(res, tt.out) {
			t.Errorf("expected: %v, got: %v", tt.out, res)
		}
	}
}
