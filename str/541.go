package str

func reverseStr(s string, k int) string {
	sChar := []byte(s)
	lastIndex := len(s) - 1
	l := 0
	r := 0
	for {
		r = l + 2*k - 1
		if r > lastIndex {
			r = lastIndex
		}
		if r-l >= k-1 { // 满k个，交换前k个
			for i := 0; i < k/2; i++ {
				sChar[l+i], sChar[l+k-1-i] = sChar[l+k-1-i], sChar[l+i] // 交换前k个
			}
		} else { // 不足k个的，全部交换
			len := r - l
			for i := 0; i <= len/2; i++ {
				sChar[l+i], sChar[r-i] = sChar[r-i], sChar[l+i]
			}
		}
		l = r + 1
		if r >= lastIndex {
			break
		}
	}
	return string(sChar)
}
