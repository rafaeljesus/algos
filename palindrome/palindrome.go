package palindrome

// str = aabaa
func IsPalindrome(str string) bool {
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}

// str=aabaa
// len = 5
// for _, i := range len(str) / 2
// i = 0
// last = 5 - 1 - 0
// str[0] = a
// strp[last] = a
// a = a? yes

// i = 1
// 5 - 1 - 1 = 3
// str[1] = a
// str[3] = a
// a = a? yes

// 1 = 2
// 5 - 1 - 2 = 2
// str[2] = b
// str[2] = b
// b = b? yes
