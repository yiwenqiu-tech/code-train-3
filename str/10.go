package str

func isMatch(s string, p string) bool {
	s = " " + s
	p = " " + p
	var f = make([][]bool, len(s))
	for i := range f {
		f[i] = make([]bool, len(p))
	}
	f[0][0] = true
	for i := 2; i < len(p); i += 2 { // 偶数位为*,例如a*可以与空字符串匹配。
		if p[i] == '*' {
			f[0][i] = true
		} else {
			break
		}
	}
	for i := 1; i < len(s); i++ {
		for j := 1; j < len(p); j++ {
			if p[j] >= 'a' && p[j] <= 'z' {
				f[i][j] = f[i-1][j-1] && (s[i] == p[j])
			} else if p[j] == '.' {
				f[i][j] = f[i-1][j-1]
			} else {
				f[i][j] = f[i-1][j] && (s[i] == p[j-1] || p[j-1] == '.')
				f[i][j] = f[i][j] || f[i][j-2]
			}
		}
	}
	return f[len(s)-1][len(p)-1]
}
