package remove_duplicates

import "fmt"

// time O(N) | Space O(1)
func RemoveDuplicates(arr []int) int {
	nextNonDup := 1
	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[nextNonDup-1] {
			arr[nextNonDup] = arr[i]
			nextNonDup++
		}
	}
	fmt.Println(arr)
	return nextNonDup
}

// time O(N) | O(1)
func CountNonDuplicates(arr []int) int {
	var count int
	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			count++
		}
	}
	return count
}

// arr: [2, 2, 2, 11]
// len(arr) = 5

// count = 0
// i = 1
// i < len(arr)? y
// arr[i] == arr[i-1]
// 2 == 2? y
// count = 1

// i = 2
// i < len(arr)? y
// arr[i] == arr[i-1]
// 2 == 2? y
// count = 2

// i = 3
// arr[i] == arr[i-1]
// 11 == 2? n
// return 2
