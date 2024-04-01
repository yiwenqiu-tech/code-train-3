package dp

type numAndCount struct {
	maxLength int
	count     int
}

func findNumberOfLIS(nums []int) int {
	n := len(nums)
	var f = make([]numAndCount, n) // 以nums[i]结尾的最长序列，其中包括最长长度以及数量。
	f[0] = numAndCount{
		maxLength: 1,
		count:     1,
	}
	for i := 1; i < n; i++ {
		maxLength := 1
		maxLengthCount := 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if f[j].maxLength+1 > maxLength {
					maxLength = f[j].maxLength + 1
					maxLengthCount = f[j].count
				} else if f[j].maxLength+1 == maxLength {
					maxLengthCount += f[j].count
				}
			}
		}
		f[i] = numAndCount{
			maxLength: maxLength,
			count:     maxLengthCount,
		}
	}
	maxLength := 0
	maxLengthCount := 0
	for i := 0; i < n; i++ {
		if f[i].maxLength > maxLength {
			maxLength = f[i].maxLength
			maxLengthCount = f[i].count
		} else if f[i].maxLength == maxLength {
			maxLengthCount += f[i].count
		}
	}
	return maxLengthCount
}
