package binary_search

func minDays(bloomDay []int, m int, k int) int {
	var minDay, maxDay int
	for _, d := range bloomDay {
		if d > maxDay {
			maxDay = d
		}
		if d < minDay {
			minDay = d
		}
	}
	left := minDay
	right := maxDay + 1 // 表示无解的情况
	for left < right {
		mid := left + (right-left)/2
		if validateForMinDays(bloomDay, m, k, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if right == maxDay+1 {
		return -1
	} else {
		return right
	}
}

func validateForMinDays(bloomDay []int, m int, k int, day int) bool {
	var accu int
	var count int
	for _, d := range bloomDay {
		if day >= d { // 满足天数
			accu++
			if accu == k {
				accu = 0
				count++
			}
		} else {
			accu = 0
		}
	}
	return count >= m
}
