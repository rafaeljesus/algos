package pattern_match

import (
	"reflect"
	"testing"
)

func TestStringMatch(t *testing.T) {
	var tests = []struct {
		text, pattern string
		found         bool
		idx           []int
	}{
		{"abaababbaba", "abba", true, []int{5, 6, 7, 8}},
		{"abada", "abada", true, []int{0, 1, 2, 3, 4}},
		{"abadaabade", "abade", true, []int{5, 6, 7, 8, 9}},
		{"abada", "abade", false, []int{}},
		{"abada", "ad*", true, []int{}},
		{"abada", "*ad", true, []int{}},
		{"abada", "a*a", true, []int{}},
	}
	for _, tt := range tests {
		found := StringMatch(tt.text, tt.pattern)
		if found != tt.found {
			t.Fatalf("text %v: expected: %v, got: %v ", tt.text, tt.found, found)
		}
	}
}

func TestStringMatchReturnIdx(t *testing.T) {
	var tests = []struct {
		text, pattern string
		found         bool
		idx           []int
	}{
		{"abaababbaba", "abba", true, []int{5, 6, 7, 8}},
		{"abada", "abada", true, []int{0, 1, 2, 3, 4}},
		{"abadaabade", "abade", true, []int{5, 6, 7, 8, 9}},
		{"abada", "abade", false, []int{}},
	}
	for _, tt := range tests {
		found, idx := StringMatchReturnIdx(tt.text, tt.pattern)
		if found != tt.found {
			t.Fatalf("expected: %v, got: %v ", tt.found, found)
		}
		if !reflect.DeepEqual(idx, tt.idx) {
			t.Fatalf("expected: %v, got: %v ", tt.idx, idx)
		}
	}
}
