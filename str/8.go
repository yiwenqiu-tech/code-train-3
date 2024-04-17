package str

import (
	"math"
	"strings"
)

func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}
	flag := 1
	skipFirst := false
	if s[0] == '-' {
		flag = -1
		skipFirst = true
	}
	if s[0] == '+' {
		flag = 1
		skipFirst = true
	}
	begin := 0
	if skipFirst {
		begin = 1
	}
	var ans = 0
	for ; begin < len(s); begin++ {
		if !(s[begin] >= '0' && s[begin] <= '9') {
			break
		}
		if flag == 1 && ans > (math.MaxInt32-int(s[begin]-'0'))/10 { // 注意ans越界，所以反过来写
			ans = math.MaxInt32
			break
		} else if flag == -1 && ans > (math.MaxInt32+1-int(s[begin]-'0'))/10 {
			flag = 1 // 这里直接赋值min了，所以flag改成1，避免后面又*flag
			ans = math.MinInt32
			break
		} else {
			ans = ans*10 + int(s[begin]-'0')
		}
	}
	ans = ans * flag
	return ans

}
