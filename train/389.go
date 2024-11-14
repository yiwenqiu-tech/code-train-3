package train

func findTheDifference(s string, t string) byte {
	var si int
	for _, c := range s {
		si = si ^ (1 << (c - 'a'))
	}
	var ti int
	for _, c := range t {
		ti = ti ^ (1 << (c - 'a'))
	}
	sti := si ^ ti // 异或
	for i := 0; i < 26; i++ {
		if sti>>i == 1 {
			return byte('a' + i)
		}
	}
	return byte(0)
}

func findTheDifference2(s string, t string) byte {
	var ans int32
	for _, c := range s {
		ans = ans ^ (c - 'a')
	}
	for _, c := range t {
		ans = ans ^ (c - 'a')
	}
	return byte(ans + 'a')
}

func findTheDifference3(s string, t string) byte {
	var rec [26]int
	for _, c := range s {
		rec[c-'a']++
	}
	for _, c := range t {
		rec[c-'a']--
		if rec[c-'a'] < 0 {
			return byte(c)
		}
	}
	return byte(0)
}
