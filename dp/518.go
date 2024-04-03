package dp

func change(amount int, coins []int) int {
	//var f = make([][]int, len(coins)+1) //  选择前n种coins时，达到amount的方案数
	//for i := range f {
	//	f[i] = make([]int, amount+1)
	//}
	//f[0][0] = 1 // 什么都不选时，总数为amount有一种方案
	//
	//for i := 1; i <= len(coins); i++ {
	//	for j := 0; j <= amount; j++ {
	//		f[i][j] = f[i-1][j]
	//		if j >= coins[i-1] {
	//			f[i][j] += f[i][j-coins[i-1]]
	//		}
	//	}
	//}
	//return f[len(coins)][amount]

	var f = make([]int, amount+1) //  选择前n种coins时，达到amount的方案数
	f[0] = 1                      // 什么都不选时，总数为amount有一种方案

	for i := 1; i <= len(coins); i++ {
		for j := coins[i-1]; j <= amount; j++ {
			f[j] += f[j-coins[i-1]]
		}
	}
	return f[amount]
}
