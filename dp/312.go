package dp

//func maxCoins(nums []int) int {
//	n := make([]int, len(nums)+2)
//	n = append(n, 1) // 数组两遍填充为0
//	n = append(n, nums...)
//	n = append(n, 1)
//
//	var f = make([][]int, len(n))
//	for i := range f {
//		f[i] = make([]int, len(n))
//		for j := range f[i] {
//			if i == j {
//				f[i][j] = n[i]
//			}
//			if i > j {
//				f[i][j] = math.MinInt32
//			}
//		}
//	}
//
//	for
//
//
//	return f[1][len(n) - 1]
//}
