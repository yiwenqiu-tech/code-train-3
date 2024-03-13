package binary_search

func shipWithinDays(weights []int, days int) int {
	var left, right int
	for _, w := range weights {
		if w > left {
			left = w
		}
		right += w
	}
	for left < right {
		mid := left + (right-left)/2
		if validateShipWithinDays(weights, days, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}

func validateShipWithinDays(weights []int, days int, shipWeight int) bool {
	var needDay = 1
	var accu int
	for _, w := range weights {
		if accu+w > shipWeight {
			accu = w
			needDay++
		} else {
			accu += w
		}
	}
	return needDay <= days
}
