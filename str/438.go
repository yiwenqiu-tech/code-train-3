package str

func findAnagrams(s string, p string) []int {
	sn := len(s)
	pn := len(p)
	if sn < pn {
		return nil
	}
	var pArray [26]int
	for i := range p {
		pArray[p[i]-'a']++
	}
	var sArray [26]int
	for i := 0; i < len(p); i++ {
		sArray[s[i]-'a']++
	}
	var ans []int
	if sArray == pArray {
		ans = append(ans, 0)
	}
	for i := pn; i < sn; i++ {
		sArray[s[i-pn]-'a']--
		sArray[s[i]-'a']++
		if sArray == pArray {
			ans = append(ans, i-pn+1)
		}
	}
	return ans
}
