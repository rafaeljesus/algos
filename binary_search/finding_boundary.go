package binary_search

func FindingBoundary(arr []bool) int {
	left := 0
	right := len(arr) - 1
	boundary_index := -1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] {
			boundary_index = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return boundary_index
}

// arr: [false, false, true, true, true]
// expected: 3
// left = 0
// right = 4
// boundary_index = -1
// left <= right? yes
// mid = (0 + 4) / 2
// mid = 2
// arr[mid] == true? yes
// boundary_index = 2
// right = mid - 1
// right = 1

// left = 0
// right = 1
// left <= right? yes
// boundary_index = 2
// mid = (0 + 1) / 2
// mid = 1 / 2 = 0
// mid = 0
// arr[mid] == true? no
// left = mid + 1
// left = 1

// left = 1
// right = 1
// boundary_index = 2
// left <= rght? yes
// mid = (1 + 1) / 2
// mid = 1
// arr[mid] == true? no
// left = mid + 1
// left = 2

// left = 2
// right = 1
// left <= right? no
// return boundary_index
