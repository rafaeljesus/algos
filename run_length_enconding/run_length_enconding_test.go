package main

import (
	"strconv"
	"testing"
)

// str "AAAAAAAAAAAAABBCCCCDD"
// out "9A4A2B4C2D"
//func RunLengthEncoding(str string) string {
//chars := make([]string, 0)
//lenght := 1
//for i := 1; i < len(str); i++ {
//if str[i] == str[i-1] && lenght < 9 {
//lenght++
//if i == len(str)-1 {
//chars = append(chars, strconv.Itoa(lenght))
//chars = append(chars, string(str[len(str)-1]))
//}
//} else {
//chars = append(chars, strconv.Itoa(lenght))
//chars = append(chars, string(str[i-1]))
//lenght = 1
//}
//}
//var res string
//for _, char := range chars {
//res += char
//}
//return res
//}

func RunLengthEncoding(str string) string {
	chars := []byte{}
	lenght := 1
	for i := 1; i < len(str); i++ {
		if str[i] != str[i-1] || lenght == 9 {
			chars = append(chars, strconv.Itoa(lenght)[0])
			chars = append(chars, str[i-1])
			lenght = 0
		}
		lenght++
	}
	chars = append(chars, strconv.Itoa(lenght)[0])
	chars = append(chars, str[len(str)-1])
	return string(chars)
}

func Test(t *testing.T) {
	var input = []struct {
		in, out string
	}{
		{"AAAAAAAAAAAAABBCCCCDD", "9A4A2B4C2D"},
		{"aA", "1a1A"},
		{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "9a9a9a9a9a9a9a9a9a9a9a9a9a9a9a9a9a9a9a9a9a9a9a1a"},
	}
	for _, tt := range input {
		res := RunLengthEncoding(tt.in)
		if res != tt.out {
			t.Errorf("expected: %s, got: %s", tt.out, res)
		}
	}
}
