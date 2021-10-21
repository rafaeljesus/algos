package remove_duplicates

// time O(N) | Space O(1)
// TODO practice test table
func RemoveDuplicates(arr []int) []int {
	i := 0
	for i < len(arr) {
		j := arr[i]
		if arr[i] != arr[j] {
			arr[j] = arr[i]
			i++
		}
	}
	return arr
}

// time O(N) | O(1)
func CountNonDuplicates(arr []int) int {
	var count = 1
	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[i-1] {
			count++
		}
	}
	return count
}
