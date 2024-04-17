package str

func strStr(haystack string, needle string) int {
	hlen := len(haystack)
	nlen := len(needle)
	if nlen > hlen {
		return -1
	}

	// 整体O(nm)
	for i := 0; i < hlen-nlen+1; i++ { // O(n)
		if haystack[i:i+nlen] == needle { // O(m)
			return i
		}
	}
	return -1
}

func strStr2(haystack string, needle string) int {
	hlen := len(haystack)
	nlen := len(needle)
	if nlen > hlen {
		return -1
	}

	d := 131
	p := 1000000007

	tmp := 1
	for i := 0; i < nlen-1; i++ {
		tmp = (tmp * d) % p // n-1次方
	}

	hHash := 0
	nHash := 0
	i := 0
	for ; i < nlen; i++ {
		hHash = (hHash*d + int(haystack[i]-'a'+1)) % p
		nHash = (nHash*d + int(needle[i]-'a'+1)) % p
	}
	i = 0
	for i <= hlen-nlen {
		if hHash == nHash && haystack[i:i+nlen] == needle {
			return i
		}
		if i+nlen >= hlen {
			break
		}
		hHash = ((hHash-int(haystack[i]-'a'+1)*tmp)*d + int(haystack[i+nlen]-'a'+1)) % p
		if hHash < 0 {
			hHash += p // 兼容负数
		}
		i++
	}
	return -1
}

func strStr3(haystack string, needle string) int {
	hn := len(haystack)
	nn := len(needle)
	// 先预处理，求前缀
	var h = make([]int64, hn+1)
	var b int64 = 131
	var p int64 = 1e9 + 7
	for i := 1; i <= hn; i++ { // 求前缀hash
		h[i] = (h[i-1]*b + (int64(haystack[i-1]) - int64('a') + 1)) % p
	}
	var nHash int64
	var base int64 = 1
	for _, c := range needle {
		nHash = (nHash*b + int64(c-'a'+1)) % p
		base = (base * b) % p // b的M次方
	}

	for l := 0; l <= hn-nn; l++ {
		r := l + nn - 1
		curHash := ((h[r+1]-h[l]*base)%p + p) % p
		if curHash == nHash {
			if haystack[l:r+1] == needle {
				return l
			}
		}
	}
	return -1
}
