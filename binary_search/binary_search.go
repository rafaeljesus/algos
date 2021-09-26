package binary_search

func BinarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := (left + right) / 2
		switch v := arr[mid]; {
		case v == target:
			return mid
		case v < target:
			left = mid + 1
		default:
			right = mid - 1
		}
	}
	return -1
}

// arr = [1, 3, 5, 7, 8]
// target = 5
// left = 0
// right = 4

// left <= right? true
// mid = 2
// v = arr[mid]
// v = 5
// v == target? true
// then return mid

// arr = [1, 2, 3, 4, 5, 6, 7]
// target = 0
// left = 0
// right = 6

// left <= right? yes
// mid = 3
// arr[mid] == 5? no
// arr[mid] < 5? no
// arr[mid] > 5? yes
// then right = mid - 1 = 2

// left <= right? yes
// mid = 0 + 2 / 2 = 0
// arr[mid] == 5? no
// arr[mid] < 5? yes
// then left = mid + 1 = 1

// left <= right? yes
// mid = 1 + 2 / 2 = 1
// arr[mid] == 5? no
// arr[mid] < 5? yes
// then left = mid + 1 = 2

// left = 2
// right = 2
// left <= right? yes
// mid = 2 + 2 / 2 = 2
// arr[mid] == 5?
// arr[mid] < 5? yes
// then left = mid + 1

// left = 3
// right = 2
// left <= right? break loop
