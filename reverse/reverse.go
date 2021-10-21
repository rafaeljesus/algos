package reverse

// Go strings are immutable and behave like read-only byte slices
// If the string only contains ASCII characters, you could also use a byte slice
// Time O(log n)? | Space O(N)
func Reverse(str string) string {
	buf := []rune(str)
	for i := 0; i < len(buf)/2; i++ {
		high := len(buf) - 1 - i
		buf[i], buf[high] = buf[high], buf[i]
	}
	return string(buf)
}

// Time O(N) | Space O(N)
func Reverse2(str string) string {
	buf := []rune(str)
	low := 0
	high := len(buf) - 1
	for low < high {
		buf[low], buf[high] = buf[high], buf[low]
		low++
		high--
	}
	return string(buf)
}

// str = golang
// tmp = str[0]
// str[0] = str[5]
// str[5] = tmp

// i = 0
// tmp = g
// str[g] = str[g]
// str[g] = str[g]

// i = 1
// tmp = str[1] = o
// str[o] = str[n]
// str[n] = str[o]
// str[gnlaog]

// i = 2
// tmp = str[2] = l
// str[l] = str[a]
// str[a] = str[l]
// str[gnalog]
