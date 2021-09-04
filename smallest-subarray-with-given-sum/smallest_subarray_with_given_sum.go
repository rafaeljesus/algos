package smallest_subarray_with_given_sum

import "math"

// time O(N+N) eg. O(N), space O(1)
func smallestSubarrayWithGivenSum(in []int, s int) int {
	var windowSum, windowStart int
	minLen := math.Inf(1)
	for windowEnd := 0; windowEnd < len(in); windowEnd++ {
		windowSum += in[windowEnd] // add next elem
		for windowSum >= s {
			minLen = math.Min(minLen, float64(windowEnd-windowStart+1))
			windowSum -= in[windowStart] // remove start elem window
			windowStart++                // move the window ahead
		}
	}
	return int(minLen)
}
