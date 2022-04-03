package main

import "fmt"

// text: abaababbaba
// pattern: abba
func FindMatch(text, pattern string) bool {
	var i, j int
	for i < len(text) {
		//fmt.Printf("text[%d]=%s pattern[%d]=%s\n", i, string(text[i]), j, string(pattern[j]))
		if text[i] != pattern[j] {
			j = 0
			continue
		}
		j++
		i++
		if j == len(pattern) {
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

	text = "abada"
	pattern = "abada"
	found = FindMatch(text, pattern)
	fmt.Println(found)

	text = "abadaabade"
	pattern = "abade"
	found = FindMatch(text, pattern)
	fmt.Println(found)
}
