package str

func validPalindrome(s string) bool {
	return checkIsValidPalindrome(s, 0, len(s)-1, 1)
}

// 递归计算，times表示可以删除多少个字符串
func checkIsValidPalindrome(s string, left int, right int, times int) bool {
	if left > right {
		return true
	}
	if s[left] == s[right] {
		return checkIsValidPalindrome(s, left+1, right-1, times)
	} else {
		if times == 0 {
			return false
		}
		return checkIsValidPalindrome(s, left, right-1, times-1) || checkIsValidPalindrome(s, left+1, right, times-1)
	}
}
