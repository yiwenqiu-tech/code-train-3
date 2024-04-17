package str

func firstUniqChar(s string) int {
	var charMap = make(map[int32]int)
	for _, c := range s {
		charMap[c]++
	}
	for index, c := range s {
		if charMap[c] == 1 {
			return index
		}
	}
	return -1
}
