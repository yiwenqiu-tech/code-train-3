package greedy

func maxProfit(prices []int) int {
	var ans int
	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1] > prices[i] {
			ans += prices[i+1] - prices[i]
		}
	}
	return ans
}
