package main

import (
	"fmt"
	"math"
)

func fullbag() {
	n := 0
	v := 0
	fmt.Scanf("%d %d", &n, &v)

	vs := make([]int, n)
	ws := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d", &vs[i], &ws[i])
	}

	//var f = make([][]int, n+1)
	//for i := range f {
	//	f[i] = make([]int, v+1)
	//}
	//f[0][0] = 0
	//
	//for i := 1; i <= n; i++ {
	//	for j := 0; j <= v; j++ {
	//		f[i][j] = f[i-1][j] // 从未选择过第i种物品
	//		if j >= vs[i-1] {
	//			f[i][j] = max(f[i][j], f[i][j-vs[i-1]]+ws[i-1]) // 选择过第i种物品
	//		}
	//	}
	//}
	//
	//var ans = math.MinInt32
	//for j := 0; j <= v; j++ {
	//	if ans < f[n][j] {
	//		ans = f[n][j]
	//	}
	//}

	var f = make([]int, v+1)
	f[0] = 0
	for i := 1; i <= n; i++ {
		for j := vs[i-1]; j <= v; j++ { // 注意这里与01背包的区别，因为这里取的是第i行，所以需要从前往后遍历，保证前面的值完成后覆盖到后面的值上。
			f[j] = max(f[j], f[j-vs[i-1]]+ws[i-1]) // 选择过第i种物品，
		}
	}

	var ans = math.MinInt32

	for j := 0; j <= v; j++ {
		if ans < f[j] {
			ans = f[j]
		}
	}
	fmt.Println(ans)
}
