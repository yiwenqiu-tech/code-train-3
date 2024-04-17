package str

func reverseOnlyLetters(s string) string { // 双指针做法，注意边界即可
	cStr := []byte(s)
	l := 0
	r := len(cStr) - 1
	for l < r {
		for l < len(cStr) && !((cStr[l] >= 'a' && cStr[l] <= 'z') || (cStr[l] >= 'A' && cStr[l] <= 'Z')) {
			l++
		}
		for r > 0 && !((cStr[r] >= 'a' && cStr[r] <= 'z') || (cStr[r] >= 'A' && cStr[r] <= 'Z')) {
			r--
		}
		if l < r {
			cStr[l], cStr[r] = cStr[r], cStr[l]
		}
		l++
		r--
	}
	return string(cStr)
}
