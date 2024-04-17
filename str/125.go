package str

var input string

func isPalindrome(s string) bool {
	input = s
	l := getNext(0)
	r := getPre(len(s) - 1)
	for l < r {
		if toLowerChar(s[l]) != toLowerChar(s[r]) {
			return false
		}
		l++
		r--
		l = getNext(l)
		r = getPre(r)
	}
	return true
}

func toLowerChar(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c - 'A' + 'a'
	}
	return c
}

func getNext(left int) int {
	for left < len(input) {
		if input[left] >= 'a' && input[left] <= 'z' || input[left] >= 'A' && input[left] <= 'Z' ||
			input[left] >= '0' && input[left] <= '9' {
			return left
		}
		left++
	}
	return left
}

func getPre(right int) int {
	for right >= 0 {
		if input[right] >= 'a' && input[right] <= 'z' || input[right] >= 'A' && input[right] <= 'Z' ||
			input[right] >= '0' && input[right] <= '9' {
			return right
		}
		right--
	}
	return right
}
