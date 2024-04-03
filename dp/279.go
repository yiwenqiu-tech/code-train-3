package dp

import "math"

func numSquares(n int) int {
	s := int(math.Sqrt(float64(n))) // n的开方，所以取数是1-s,目标是N
	//
	//var f = make([][]int, s+1) // 利用前i种数，达到j所达到的最小数量
	//for i := range f {
	//	f[i] = make([]int, n+1)
	//	for j := range f[i] {
	//		f[i][j] = math.MaxInt32
	//	}
	//}
	//f[0][0] = 0
	//
	//for i := 1; i <= s; i++ {
	//	for j := 0; j <= n; j++ {
	//		f[i][j] = f[i-1][j]
	//		if j >= i * i {
	//			f[i][j] = min(f[i][j], f[i][j- i * i]+1)
	//		}
	//	}
	//}
	//return f[s][n]

	f := make([]int, n+1)
	for j := range f {
		f[j] = math.MaxInt32
	}
	f[0] = 0

	for i := 1; i <= s; i++ {
		for j := i * i; j <= n; j++ {
			f[j] = min(f[j], f[j-i*i]+1)
		}
	}
	return f[n]

}
