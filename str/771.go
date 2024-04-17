package str

func numJewelsInStones(jewels string, stones string) int {
	jewelMap := make(map[int32]struct{})
	for _, c := range jewels {
		jewelMap[c] = struct{}{}
	}

	var ans = 0
	for _, c := range stones {
		if _, exist := jewelMap[c]; exist {
			ans++
		}
	}
	return ans
}
