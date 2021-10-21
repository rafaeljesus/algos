package twosum

import "sort"

// O(n) time | O(n) space
func CheckPairWithSum1(arr []int, target int) (bool, []int) {
	nums := make(map[int]struct{})
	for _, num := range arr {
		potentialMatch := target - num
		if _, ok := nums[potentialMatch]; ok {
			return true, []int{potentialMatch, num}
		}
		nums[num] = struct{}{}
	}
	return false, []int{}
}

// Time: O(nlog(n))
// Space: O(1)
func CheckPairWithSum2(arr []int, target int) (bool, []int) {
	sort.Ints(arr)
	low := 0
	high := len(arr) - 1
	for low < high {
		sum := arr[low] + arr[high]
		if sum == target {
			return true, []int{arr[low], arr[high]}
		} else if sum < target {
			low++
		} else {
			high--
		}
	}
	return false, []int{}
}
