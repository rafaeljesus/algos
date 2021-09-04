package max_sum_subarray_sizek

import "testing"

func Test(t *testing.T) {
	in := []int{2, 1, 5, 1, 3, 2}
	k := 3
	out := maxSumSubarraySizek(in, k)
	expected := 9
	if out != expected {
		t.Errorf("expected: %v,  got: %v", expected, out)
	}
}
