package dp

var nums416 []int
var target int
var can416 bool

func canPartition(nums []int) bool {
	var sum = 0
	for i := range nums {
		sum += nums[i]
	}
	if sum%2 == 1 { // 奇数，直接返回false
		return false
	}
	nums416 = nums
	target = sum / 2

	//var f = make([][]bool, len(nums)+1)
	//for i := range f {
	//	f[i] = make([]bool, target+1)
	//}
	//f[0][0] = true // 不选数，和为0为true
	//for i := 1; i <= len(nums); i++ {
	//	for j := 0; j <= target; j++ {
	//		f[i][j] = f[i-1][j]
	//		if j >= nums[i-1] {
	//			f[i][j] = f[i][j] || f[i-1][j-nums[i-1]] // 如果f[i-1][j-nums[i-1]]为true，表示+nums[i-1]可以到j
	//		}
	//	}
	//}
	//
	//for i := 1; i <= len(nums); i++ {
	//	if f[i][target] {
	//		return true
	//	}
	//}
	//return false

	// 优化版，从上面可知，f[i][j]初始化时为f[i-1][j]，然后与f[i-1][j-nums[i-1]]求或。主要j需要逆向求。
	f := make([]bool, target+1)
	f[0] = true // 不选数，和为0为true
	for i := 1; i <= len(nums); i++ {
		for j := target; j >= nums[i-1]; j-- {
			f[j] = f[j] || f[j-nums[i-1]] // 如果f[i-1][j-nums[i-1]]为true，表示+nums[i-1]可以到j
		}
	}
	return f[target]
}

func dfsForCanPartition(i int, cur int) { // 状态为i:=遍历到第几个数，cur:当前值多少
	if cur > target {
		return
	}
	if cur == target {
		can416 = true
		return
	}
	if i >= len(nums416) {
		return
	}
	dfsForCanPartition(i+1, cur)
	cur += nums416[i]
	dfsForCanPartition(i+1, cur)
	cur -= nums416[i]
}
