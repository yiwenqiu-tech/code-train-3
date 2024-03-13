package binary_search

func mySqrt(x int) int {
	left := 0
	right := x
	for left < right {
		mid := left + (right-left+1)/2
		if mid*mid <= x { // 找前继
			left = mid
		} else {
			right = mid - 1
		}
	}
	return right
	//return int(doubleForMySqrt(x))
}

func doubleForMySqrt(x int) float64 { // 实数二分
	var left float64 = 0
	var right = float64(x)

	for right-left > 0.00001 {
		mid := left + (right-left)/2
		if mid*mid <= float64(x) {
			left = mid
		} else {
			right = mid
		}
	}
	return right
}
