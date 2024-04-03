package main

import (
	"fmt"
	"math"
)

var vs []int
var ws []int
var v int
var ans int

func bag01() {
	n := 0
	v = 0
	ans = 0
	fmt.Scanf("%d %d", &n, &v)

	vs = make([]int, n)
	ws = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d", &vs[i], &ws[i])
	}

	//var f = make([][]int, n+1) // 表示前i个 j体积下获取到的最大价值
	//for i := range f {
	//	f[i] = make([]int, v+1)
	//	for j := range f[i] {
	//		f[i][j] = math.MinInt32
	//	}
	//}
	//f[0][0] = 0
	//
	//for i := 1; i <= n; i++ {
	//	for j := 0; j <= v; j++ {
	//		f[i][j] = f[i-1][j] // 当前i
	//		if j >= vs[i-1] {
	//			f[i][j] = max(f[i-1][j-vs[i-1]]+ws[i-1], f[i][j])
	//		}
	//	}
	//}
	//var ans = math.MinInt32
	//for j := 0; j <= v; j++ {
	//	if f[n][j] > ans {
	//		ans = f[n][j]
	//	}
	//}

	var f = make([]int, v+1) // 优化空间版本，从上面可以看到，每次f[i][j]都拿f[i-1][j]来赋值，那其实这里可以共用f[j]来减少空间
	f[0] = 0
	for i := 1; i <= n; i++ {
		for j := v; j >= vs[i-1]; j-- { // 从后往前，避免重复覆盖值
			f[j] = max(f[j-vs[i-1]]+ws[i-1], f[j])
		}
	}

	var ans = math.MinInt32
	for j := 0; j <= v; j++ {
		if f[j] > ans {
			ans = f[j]
		}
	}

	fmt.Println(ans)

}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 搜索
func dfsFor01Bag(i int, cur int, weight int) {
	if cur > v { // 容积超了
		return
	}
	if i == len(ws) { // 遍历到最后一个了
		if weight > ans {
			ans = weight
		}
		return
	}
	dfsFor01Bag(i+1, cur, weight) // 不要
	cur += vs[i]
	weight += ws[i]
	dfsFor01Bag(i+1, cur, weight) // 要
	cur -= vs[i]
	weight -= ws[i]
}
