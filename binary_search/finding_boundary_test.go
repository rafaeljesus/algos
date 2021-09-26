package binary_search

import (
	"reflect"
	"testing"
)

func TestFindingBoundary(t *testing.T) {
	var input = []struct {
		in  []bool
		out int
	}{
		{[]bool{false, false, true, true, true}, 2},
	}
	for _, tt := range input {
		res := FindingBoundary(tt.in)
		if !reflect.DeepEqual(res, tt.out) {
			t.Errorf("expected %v, got: %v", tt.out, res)
		}
	}
}
