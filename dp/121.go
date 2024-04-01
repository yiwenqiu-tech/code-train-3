package dp

import "math"

func maxProfit(prices []int) int {
	var min = math.MaxInt32
	var max = 0
	for _, p := range prices {
		if p < min {
			min = p
		} else {
			if p-min > max {
				max = p - min
			}
		}
	}
	return max
}
