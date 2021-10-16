package palindrome

import "testing"

func TestPalindrome(t *testing.T) {
	var input = []struct {
		in  string
		out bool
	}{
		{"aabaa", true},
		{"aaabaa", false},
	}
	for _, tt := range input {
		if IsPalindrome(tt.in) != tt.out {
			t.Errorf("should be palindrome")
		}
	}
}
