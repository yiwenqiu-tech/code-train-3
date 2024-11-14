package train

import "math"

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	existCharMap := make(map[uint8]struct{})
	maxLength := math.MinInt
	r := -1

	for i := 0; i < len(s); i++ {
		if i > 0 {
			delete(existCharMap, s[i-1]) // i > 0,循环到这里，相当于已经出现重复了，从头元素开始删除
		}
		for r+1 < len(s) && !inCharMap(s[r+1], existCharMap) {
			existCharMap[s[r+1]] = struct{}{}
			r++
		}
		length := r - i + 1
		if length > maxLength {
			maxLength = length
		}
	}
	return maxLength
}

func inCharMap(c uint8, existCharMap map[uint8]struct{}) bool {
	if _, exist := existCharMap[c]; exist {
		return true
	}
	return false
}
