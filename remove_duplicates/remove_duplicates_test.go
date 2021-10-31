package remove_duplicates

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	var input = []struct {
		in  []int
		out int
	}{
		// Explanation:
		// The first four elements after removing the duplicates will be [2, 3, 6, 9].
		{[]int{2, 3, 3, 3, 6, 9, 9}, 4},
		{[]int{2, 2, 2, 11}, 2},
	}
	for _, tt := range input {
		res := RemoveDuplicates(tt.in)
		if !reflect.DeepEqual(res, tt.out) {
			t.Errorf("expected %v, got: %v", tt.out, res)
		}
	}
}

func TestCountNonDuplicates(t *testing.T) {
	var input = []struct {
		in  []int
		out int
	}{
		// Explanation:
		// The first four elements after removing the duplicates will be [2, 3, 6, 9].
		{[]int{2, 3, 3, 3, 6, 9, 9}, 3},
		{[]int{2, 2, 2, 11}, 2},
	}
	for _, tt := range input {
		res := CountNonDuplicates(tt.in)
		if res != tt.out {
			t.Errorf("expected %v, got: %v", tt.out, res)
		}
	}
}
