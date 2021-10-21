package cyclic_sort

import (
	"reflect"
	"testing"
)

func TestCyclicSort(t *testing.T) {
	var input = []struct {
		in, out []int
	}{
		{[]int{1, 5, 6, 4, 3, 2}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{3, 1, 5, 4, 0}, []int{1, 2, 3, 4, 5}},
	}
	for _, tt := range input {
		res := CyclicSort(tt.in)
		if !reflect.DeepEqual(res, tt.out) {
			t.Errorf("expected: %v, got: %v", tt.out, res)
		}
	}
}
