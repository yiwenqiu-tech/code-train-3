package train

import "strings"

func wordPattern(pattern string, s string) bool {
	sslice := strings.Split(s, " ")
	if len(sslice) != len(pattern) {
		return false
	}
	var patternSMap = make(map[int32]string)
	var sPatternMap = make(map[string]int32)
	for index, c := range pattern {
		if tmp, exist := patternSMap[c]; exist {
			if sslice[index] != tmp {
				return false
			}
		} else {
			if _, existInSPattern := sPatternMap[sslice[index]]; existInSPattern {
				return false
			}
			patternSMap[c] = sslice[index]
			sPatternMap[sslice[index]] = c

		}
	}
	return true
}
