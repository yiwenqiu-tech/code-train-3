package str

import "strings"

func lengthOfLastWord(s string) int {
	ss := strings.Split(strings.TrimSpace(s), " ")
	return len(ss[len(ss)-1])

	// TODO 不用库函数
}
