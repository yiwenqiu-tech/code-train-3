package _map

import "sort"

func groupAnagrams(strs []string) [][]string {
	var strMap = make(map[string][]string)
	for _, str := range strs {
		bytes := []byte(str)
		bs := byteSlice(bytes)
		sort.Sort(bs)
		sortStr := string(bs)
		strMap[sortStr] = append(strMap[sortStr], str)
	}
	var res [][]string
	for _, value := range strMap {
		res = append(res, value)
	}
	return res
}

type byteSlice []byte

func (x byteSlice) Len() int           { return len(x) }
func (x byteSlice) Less(i, j int) bool { return x[i] < x[j] }
func (x byteSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
