package presum

func numberOfSubarrays(nums []int, k int) int {
	var tmp = make([]int, len(nums))
	for index, num := range nums {
		tmp[index] = num % 2 // 转换为奇偶数组，偶数时为0，奇数时为1
	}
	var s = make([]int, len(nums)+1) // 前缀和数组
	s[0] = 0
	for index, num := range tmp {
		s[index+1] = s[index] + num
	}
	var diffMap = make(map[int]int)
	var res int
	for _, num := range s {
		if count, exist := diffMap[num-k]; exist {
			res += count
		}
		diffMap[num]++
	}
	return res
}
