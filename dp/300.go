package dp

import "fmt"

var maxLen int

func lengthOfLIS2(nums []int) int {
	maxLen = 0
	dfsForLIS(nums, 0, nil)
	return maxLen
}

func dfsForLIS(nums []int, i int, chosen []int) {
	if i == len(nums) {
		if len(chosen) > maxLen {
			maxLen = len(chosen)
		}
		return
	}
	dfsForLIS(nums, i+1, chosen)

	// 还没选，或者当前数比已选的最后一个数大
	if len(chosen) == 0 || nums[i] > chosen[len(chosen)-1] {
		chosen = append(chosen, nums[i])

		dfsForLIS(nums, i+1, chosen)

		chosen = chosen[:len(chosen)-1]
	}

}

func lengthOfLIS(nums []int) int {
	l := len(nums)
	var f = make([]int, l)
	var pre = make([]int, l)
	for i := 0; i < l; i++ {
		f[i] = 1 // 每个元素至少有他一个元素的路径
		pre[i] = -1
	}
	for i := 1; i < l; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if f[j]+1 > f[i] {
					f[i] = f[j] + 1
					pre[i] = j // 记录从j转移到i
				}
			}
		}
	}
	var ans int
	var end int
	for i := 0; i < l; i++ {
		if f[i] > ans {
			ans = f[i]
			end = i
		}
	}
	print(pre, nums, end)
	return ans
}

func print(pre []int, nums []int, i int) {
	if pre[i] != -1 {
		print(pre, nums, pre[i])
	}
	fmt.Println(nums[i])
}
