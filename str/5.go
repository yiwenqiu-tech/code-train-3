package str

func longestPalindrome(s string) string {
	var ansLen int
	var ansStart int
	// 奇数长度的回文串
	for i := 0; i < len(s); i++ {
		left := i - 1
		right := i + 1
		// left + 1 ~ i - i ~ right -1
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
		}
		if (right-1-i)*2+1 > ansLen {
			ansLen = (right-1-i)*2 + 1
			ansStart = left + 1
		}
	}
	// 偶数长度的回文串
	for i := 1; i < len(s); i++ {
		left := i - 1
		right := i
		// left + 1 ~ i - i ~ right
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
		}
		if (right-i)*2 > ansLen {
			ansLen = (right - i) * 2
			ansStart = left + 1
		}
	}
	return s[ansStart : ansStart+ansLen]

	// TODO RabinKarp hash + 二分答案
}
