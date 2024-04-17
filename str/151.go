package str

import (
	"regexp"
	"strings"
)

func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	reg := regexp.MustCompile(`\s+`)
	sSlice := reg.Split(s, -1)
	for i := 0; i < len(sSlice)/2; i++ {
		sSlice[i], sSlice[len(sSlice)-1-i] = sSlice[len(sSlice)-1-i], sSlice[i]
	}
	return strings.Join(sSlice, " ")

	// TODO O(1)空间做法？
}
