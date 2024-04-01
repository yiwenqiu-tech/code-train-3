package greedy

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	var ans int
	j := len(g) - 1 // 吃的最多的小孩
	for i := len(s) - 1; i >= 0 && j >= 0; i-- {
		for j >= 0 && s[i] < g[j] { // 当前不满足j的分量，直接j--
			j--
		}
		if j >= 0 {
			ans++
			j--
		}
	}
	return ans
}
