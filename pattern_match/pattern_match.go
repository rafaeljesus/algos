package pattern_match

// Time: O(n) | Space: O(1)
func FindMatch(text, pattern string) bool {
	var i, j int
	for i < len(text) {
		if text[i] != pattern[j] {
			j = 0
			continue
		}
		j++
		if j == len(pattern) {
			return true
		}
		i++
	}
	return false
}

// Time: O(n) | Space: O(n)
func FindMatchAndReturnIdx(text, pattern string) (bool, []int) {
	idx := []int{}
	var i, j int
	for i < len(text) {
		if text[i] != pattern[j] {
			j = 0
			idx = []int{}
			continue
		}
		idx = append(idx, i)
		j++
		if j == len(pattern) {
			return true, idx
		}
		i++
	}
	return false, []int{}
}
