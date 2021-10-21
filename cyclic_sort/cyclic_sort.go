package cyclic_sort

// in: [1, 5, 6, 4, 3, 2]
// out: [1, 2, 3, 4, 5, 6]
// Time|Space O(N)
func CyclicSort(nums []int) []int {
	var i int
	for i < len(nums) {
		j := nums[i] - 1
		if nums[i] != nums[j] {
			nums[i], nums[j] = nums[j], nums[i] // swap
		} else {
			i++
		}
	}
	return nums
}

// nums = [1, 5, 6, 4, 3, 2]
// lenNums = 6
// i = 0
// i < lenNums? y
// j = 1 - 0
// if 1 != 0? y
// nums[0] 1 = nums[1] 5
// nums[5, 1, 6, 4, 3, 2]

// nums[5, 1, 6, 4, 3, 2]
// i = 0
// i < lenNums? y
// j = 4 = (5 -1 )
// if 5 != 3? y
// nums[0] 5 = nums[4] 3
// nums [3, 1, 6, 4, 5, 2]

// nums [3, 1, 6, 4, 5, 2]
// i = 0
// i < lenNums? y
// j = 2 (3 -1)
// if 3 != 6? y
// nums[0] 3 = nums[2] 6
// nums [6, 1, 3, 4, 5, 2]

// nums [6, 1, 3, 4, 5, 2]
// i = 0
// j = 5 (6 -1)
// if 6 != 2? y
// nums[0] 6 = nums[5] 2
// nums = [2, 1, 3, 4, 5, 6]

// nums = [2, 1, 3, 4, 5, 6]
// i = 0
// j = 1 = (2 - 1)
// if 2 != 1? y
// nums[0] 2 = nums[1] 1
// nums = [1, 2, 3, 4, 5, 6]

// nums = [1, 2, 3, 4, 5, 6]
// i = 0
// j = 0 (1 - 1)
// if 0 != 0? n
// then i = 1

// nums = [1, 2, 3, 4, 5, 6]
// i = 1
// j = 1 (2 - 1)
// if 2 != 2 n
// then i = 2

// nums = [1, 2, 3, 4, 5, 6]
// i = 2
// j = 2 (3 - 1)
// if 3 != 3 n
// then i = 3

// keeps until 1 < lenNums is false
