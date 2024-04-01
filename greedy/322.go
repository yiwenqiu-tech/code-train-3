package greedy

import (
	"math"
)

var min = 10001
var inputCoins []int

func coinChange(coins []int, amount int) int {
	min = 10001
	inputCoins = coins
	bfsForCoinChange(amount, 0)
	if min == 10001 {
		return -1
	}
	return min
}

func bfsForCoinChange(amount int, used int) {
	if amount < 0 {
		return
	}
	if amount == 0 {
		if used < min {
			min = used
		}
	}
	for _, c := range inputCoins {
		bfsForCoinChange(amount-c, used+1)
	}
}

func coinChange2(coins []int, amount int) int {
	amountCoin := make([]int, amount+1)
	amountCoin[0] = 0

	for i := 1; i <= amount; i++ {
		amountCoin[i] = math.MaxInt32
		for _, c := range coins {
			if i-c >= 0 {
				amountCoin[i] = int(math.Min(float64(amountCoin[i]), float64(amountCoin[i-c]+1)))
			}
		}
	}
	if amountCoin[amount] == math.MaxInt32 {
		return -1
	}
	return amountCoin[amount]
}

var amountAns []int

func coinChange3(coins []int, amount int) int {
	amountAns = make([]int, amount+1)
	for i := range amountAns {
		amountAns[i] = -1
	}
	amountAns[0] = 0
	ans := dfsForCoinChange(coins, amount)
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}

func dfsForCoinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return math.MaxInt32
	}
	if amountAns[amount] != -1 {
		return amountAns[amount]
	}
	amountAns[amount] = math.MaxInt32
	for _, c := range coins {
		amountAns[amount] = int(math.Min(float64(amountAns[amount]), float64(dfsForCoinChange(coins, amount-c)+1)))
	}
	return amountAns[amount]
}
