package main

import "fmt"

// text: abaababbaba
// pattern: abba
func FindMatch(text, pattern string) bool {
	var count int
	j := 0
	for i := 0; i < len(text); i++ {
		if text[i] != pattern[j] {
			j = 0
			continue
		}
		count++
		if count == len(pattern) {
			return true
		}
	}
	return false
}

func main() {
	text := "abaababbaba"
	pattern := "abba"
	found := FindMatch(text, pattern)
	fmt.Println(found)

	text = "abada"
	pattern = "abade"
	found = FindMatch(text, pattern)
	fmt.Println(found)

	text = "abadaabade"
	pattern = "abade"
	found = FindMatch(text, pattern)
	fmt.Println(found)
}
