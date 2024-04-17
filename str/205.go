package str

func isIsomorphic(s string, t string) bool {
	var sMap = make(map[byte]byte)
	var tMap = make(map[byte]byte)
	for i := range s {
		if c, exist := sMap[s[i]]; exist { // s的字符之前是否出现过，若出现过，映射里的字符是否等于t里的
			if c != t[i] {
				return false
			}
		} else { // s的字符之前没出现过，那该字符在t里是否出现过，若出现过，表明存在不匹配的映射，直接返回false
			if _, tExist := tMap[t[i]]; tExist {
				return false
			}
			sMap[s[i]] = t[i]
			tMap[t[i]] = s[i]
		}
	}
	return true
}
