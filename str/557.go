package str

import (
	"strings"
)

func reverseWords557(s string) string {
	sSlice := strings.Split(s, " ")
	for i := 0; i < len(sSlice); i++ {
		cStr := []byte(sSlice[i])
		for j := 0; j < len(cStr)/2; j++ {
			cStr[j], cStr[len(cStr)-1-j] = cStr[len(cStr)-1-j], cStr[j]
		}
		sSlice[i] = string(cStr)
	}
	return strings.Join(sSlice, " ")

	// TODO 原生做法
}
