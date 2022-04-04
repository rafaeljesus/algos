package word_pattern

import "testing"

func TestWordPattern(t *testing.T) {
	var tests = []struct {
		pattern, text string
		found         bool
	}{
		{"abba", "dog cat cat dog", true},
		{"abba", "dog cat cat fish", false},
		{"abba", "dog dog dog dog", false},
		{"aaaa", "dog cat cat fish", false},
		{"aaa", "aa aa aa aa", false},
	}
	for _, tt := range tests {
		found := WordPattern(tt.pattern, tt.text)
		if found != tt.found {
			t.Errorf("expected: %v, got: %v", tt.found, found)
		}
	}
}
