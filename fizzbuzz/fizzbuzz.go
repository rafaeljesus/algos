package fizzbuzz

import "strconv"

/*
Given an integer n, return a string array answer (1-indexed) where:
answer[i] == "FizzBuzz" if i is divisible by 3 and 5.
answer[i] == "Fizz" if i is divisible by 3.
answer[i] == "Buzz" if i is divisible by 5.
answer[i] == i if non of the above conditions are true.

Example 1:
Input: n = 3
Output: ["1","2","Fizz"]

Example 2:
Input: n = 5
Output: ["1","2","Fizz","4","Buzz"]
*/
func Fizzbuzz(num int) []string {
	fizzbuzzMap := map[int]string{3: "Fizz", 5: "Buzz"}
	res := make([]string, 0)
	for i := 1; i <= num; i++ {
		var str string
		for key, val := range fizzbuzzMap {
			if i%key == 0 {
				str += val
			}
		}
		if str == "" {
			str = strconv.Itoa(i)
		}
		res = append(res, str)
	}
	return res
}

// n = 3

// i = 1
// str = ""
// i%val == 0
// i%3 == 0? n
// str = "1"
// res = ["1"]

// i = 2
// key = 3, val = Fizz
// i % key == 0? n
// key = 5, val = Buzz
// i % key == 0? n
// str = "2"
// res = ["1", "2"]

// i = 3
// key = 3, val = Fizz
// i % key == 0
// 3 % 3 == 0? y
// str = "Fizz"
// key = 5, val = Buzz
// 3 % 5 == 0? n
// res = ["1", "2", "Fizz"]
