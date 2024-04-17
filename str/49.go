package str

func groupAnagrams(strs []string) [][]string {
	//var strMap = make(map[string][]string)
	//for i := range strs {
	//	tmp := []byte(strs[i])
	//	sort.Slice(tmp, func(i, j int) bool {
	//		return tmp[i] < tmp[j]
	//	})
	//	strMap[string(tmp)] = append(strMap[string(tmp)], strs[i])
	//}
	//var ans [][]string
	//for _, v := range strMap {
	//	ans = append(ans, v)
	//}
	//return ans

	var strMap = make(map[[26]int][]string)
	for i := range strs {
		var label [26]int
		for j := range strs[i] {
			label[strs[i][j]-'a']++
		}
		strMap[label] = append(strMap[label], strs[i])
	}
	var ans [][]string
	for _, v := range strMap {
		ans = append(ans, v)
	}
	return ans
}
