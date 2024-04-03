package main

import "fmt"

func main() {
	var k int
	fmt.Scanf("%d", &k)

	var al int
	var ac int
	var bl int
	var bc int
	fmt.Scanf("%d %d %d %d", &al, &ac, &bl, &bc)

	var totalList = make([]int, ac+bc)
	for i := 0; i < ac; i++ {
		totalList[i] = al
	}
	for i := ac; i < ac+bc; i++ {
		totalList[i] = bl
	}

	//var f = make([][]int, ac+bc+1) // 前i个物品达到j的方案数
	//for i := range f {
	//	f[i] = make([]int, k+1)
	//}
	//
	//f[0][0] = 1
	//for i := 1; i <= ac+bc; i++ {
	//	for j := 0; j <= k; j++ {
	//		f[i][j] = f[i-1][j]
	//		if j >= totalList[i-1] {
	//			f[i][j] = (f[i][j] + f[i-1][j-totalList[i-1]]) % 1000000007
	//
	//		}
	//	}
	//}

	f := make([]int, k+1)
	f[0] = 1
	for i := 1; i <= ac+bc; i++ {
		for j := k; j >= totalList[i-1]; j-- {
			f[j] = (f[j] + f[j-totalList[i-1]]) % 1000000007
		}
	}

	fmt.Println(f[k])
	return
}
