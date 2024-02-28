package presum

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	var ans int
	for i := 0; i < len(matrix); i++ {
		var sum = make([]int, len(matrix[0]))
		for j := i; j < len(matrix); j++ {
			for k := 0; k < len(matrix[0]); k++ { // 遍历列
				sum[k] = sum[k] + matrix[j][k]
			}
			ans += matrixSumMatchTargetCount(sum, target)
		}
	}
	return ans
}

func matrixSumMatchTargetCount(sum []int, target int) int {
	var preSum = make([]int, len(sum)+1)
	for i := 0; i < len(sum); i++ {
		preSum[i+1] = preSum[i] + sum[i]
	}

	var res int
	var countMap = make(map[int]int)
	for i := 0; i < len(preSum); i++ {
		find := preSum[i] - target
		if count, exist := countMap[find]; exist {
			res += count
		}
		countMap[preSum[i]]++
	}

	return res

}
