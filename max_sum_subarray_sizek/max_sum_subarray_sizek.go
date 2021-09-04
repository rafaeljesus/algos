package max_sum_subarray_sizek

import "math"

func maxSumSubarraySizek(in []int, k int) int {
	var maxSum, windowSum, windowStart int
	for windowEnd := 0; windowEnd < len(in); windowEnd++ {
		windowSum += in[windowEnd] // add next element
		if windowEnd >= k-1 {
			maxSum = int(math.Max(float64(maxSum), float64(windowSum)))
			windowSum -= in[windowStart] // remove start elem window
			windowStart++                // move the window ahead
		}
	}
	return maxSum
}

// test case
//Input: [2, 1, 5, 1, 3, 2], k=3
//windowEnd = 0
//windowSum = 2

//windowEnd = 1
//windowSum = 2 + 1 = 3

//windowEnd = 2
//windowSum = 3 + 5 = 8

//2 >= 2 == true // [2, 1, 5]
//maxSum = 8 (max of maxSum and windowSum)
//windowEnd -= 2 (first element of the window)
//windowStart = 1 (the start of the window)
//// [1, 5, 1]
//windowEnd = 3
//windowSum = 6 + 1 = 7

//windowEnd = 3
//windowSum = 6 + 1 = 7
//maxSum = 8
