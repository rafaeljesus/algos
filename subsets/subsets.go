package subsets

// in: [1, 3]
// out: [[], [1], [3], [1, 3]]
func FindSubsets(nums []int) [][]int {
	subsets := [][]int{[]int{}}
	for _, num := range nums {
		n := len(subsets)
		for i := 0; i < n; i++ {
			set := subsets[i]
			set = append(set, num)
			subsets = append(subsets, set)
		}
	}
	return subsets
}

// nums [1, 3]
// subsets [[]]
// num = 1
// n = 1
// i = 0; 0 < 1? y
// set = []
// set = [1]
// subsets = [[], [1]]
// i = 1; 1 < 1? n

// num = 3
// n = 2
// i = 0; 0 < 2? y
// set = []
// set = [3]
// subsets = [[], [1], [3]]

// i = 1; 1 < 2? y
// set = [1]
// set = [1, 3]
// subsets = [[], [1], [3], [1, 3]]

// i = 2; 2 < 2? n
// return [[], [1], [3], [1, 3]]
