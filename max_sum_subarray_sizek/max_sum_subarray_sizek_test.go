package max_sum_subarray_sizek

import (
	"testing"
)

func Test(t *testing.T) {
	in := []int{2, 1, 5, 1, 3, 2}
	k := 3
	out := maxSumSubarraySizek(in, k)
	//out := test(in, k)
	expected := 9
	if out != expected {
		t.Errorf("expected: %v,  got: %v", expected, out)
	}
}

func test(arr []int, k int) int {
	var maxSum, windowSum, windowStart int
	for windonEnd := 0; windonEnd < len(arr); windonEnd++ {
		windowSum += arr[windowEnd]
		if windowEnd >= k-1 {
			maxSum = max(maxSum, windowSum)
			windowSum -= arr[windowStart]
			windowStart++
		}
	}
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
