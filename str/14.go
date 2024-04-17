package str

import "strings"

func longestCommonPrefix(strs []string) string {
	var prefix strings.Builder
	var point = 0
	var end bool
	for {
		var c byte
		for _, str := range strs {
			if len(str) <= point { // 到底了
				end = true
				break
			}
			if c == 0 {
				c = str[point] // 取第一个字符串的字符来做标准
			} else {
				if c != str[point] { // 不等于第一个字符串的字符，终止
					end = true
					break
				}
			}
		}
		if end {
			break
		}
		prefix.WriteByte(c)
		point++

	}
	return prefix.String()
}
