package word_pattern

import (
	"strings"
)

// Time: O(m) | Space O(m)
// m is the number of uniq chars in pattern
func WordPattern(pattern, text string) bool {
	patternMap := map[byte]string{}
	wordMap := map[string]byte{}
	words := strings.Split(text, " ")
	if len(pattern) != len(words) {
		return false
	}
	for i := 0; i < len(pattern); i++ {
		if word, ok := patternMap[pattern[i]]; ok {
			if word != words[i] {
				return false
			}
		}
		if p, ok := wordMap[words[i]]; ok {
			if p != pattern[i] {
				return false
			}
		}
		// patternMap{a: dog}
		patternMap[pattern[i]] = words[i]
		// wordMap{dog: a}
		wordMap[words[i]] = pattern[i]
	}
	return true
}
