package binary_search

// arr: [30, 40, 50, 10, 20], target: 3
// expected: 3
func FindMinimunInRotatedSortedArray(arr []int) int {
	left := 0
	right := len(arr) - 1
	boundaryIndex := -1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] < arr[len(arr)-1] {
			boundaryIndex = mid
			right = mid + 1
		} else {
			left = mid - 1
		}
	}
	return boundaryIndex
}
