package str

func isAnagram(s string, t string) bool {
	sn := len(s)
	tn := len(t)
	if sn != tn {
		return false
	}
	//var sMap = make(map[byte]int)
	//var tMap = make(map[byte]int)
	//for i := 0; i < sn; i++ {
	//	sMap[s[i]]++
	//	tMap[t[i]]++
	//}
	//if len(sMap) != len(tMap) {
	//	return false
	//}
	//for k, v := range sMap {
	//	if tv, exist := tMap[k]; exist {
	//		if tv != v {
	//			return false
	//		}
	//	} else {
	//		return false
	//	}
	//}
	//return true

	var f = make([]int, 26)
	for i := 0; i < sn; i++ {
		f[s[i]-'a']++
		f[t[i]-'a']--
	}
	for i := range f {
		if f[i] != 0 {
			return false
		}
	}
	return true
}
