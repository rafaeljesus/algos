package twosum

import "sort"

func CheckPairWithSum(arr []int, target int) (bool, []int) {
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

func CheckPairWithSumOptimal(arr []int, target int) (bool, []int) {
	//m := make(map[int]int)
}
