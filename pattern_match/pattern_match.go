package pattern_match

// Time: O(nm) | Space: O(1)
func StringMatch(text, pattern string) bool {
	n := len(text)
	m := len(pattern)
	for i := 0; i <= n-m; i++ {
		j := 0
		for j < m && (text[i+j] == pattern[j] || pattern[j] == '*') {
			j++
		}
		if j == m {
			return true
		}
	}
	return false
}

// Time: O(nm) | Space: O(n)
func StringMatchReturnIdx(text, pattern string) (bool, []int) {
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
