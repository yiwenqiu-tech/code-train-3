package train

func minWindow(s string, t string) string {
	l := 0
	r := 0
	targetMap := make(map[uint8]int)
	for index := range t {
		targetMap[t[index]]++
	} // targetMap
	var ans string
	tmpMap := make(map[uint8]int)
	for {
		if match(tmpMap, targetMap) { // match了，移动左端点
			if len(ans) == 0 || r-l < len(ans) { // 为空，或者比ans小，更新ans
				ans = s[l:r]
			}
			// 移动左端点
			tmpMap[s[l]]--
			l++
		} else {
			if r >= len(s) {
				break
			}
			if _, exist := targetMap[s[r]]; exist {
				tmpMap[s[r]]++
			}
			r++
		}
	}
	return ans
}

func match(srcMap, targetMap map[uint8]int) bool {
	for k, v := range targetMap {
		if srcMap[k] < v {
			return false
		}
	}
	return true
}
