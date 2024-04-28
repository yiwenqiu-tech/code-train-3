package str

func numDistinct(s string, t string) int {
	s = " " + s
	t = " " + t
	var f = make([][]int, len(s))
	for i := range f {
		f[i] = make([]int, len(t))
	}
	for i := 0; i < len(s); i++ {
		f[i][0] = 1
	}
	for i := 1; i < len(s); i++ {
		for j := 1; j < len(t); j++ {
			f[i][j] = f[i-1][j] // 第i个字符不使用
			if s[i] == t[j] {   // 当s的第i个字符与s的第j个字符相等时，方案数加上f[i-1][j-1]
				f[i][j] += f[i-1][j-1]
			}
		}
	}
	return f[len(s)-1][len(t)-1]
}
