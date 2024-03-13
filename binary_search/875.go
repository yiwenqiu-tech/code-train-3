package binary_search

import "math"

func minEatingSpeed(piles []int, h int) int {
	var left = 1
	var right int
	for _, p := range piles {
		if p > right {
			right = p
		}
	}
	for left < right {
		mid := left + (right-left)/2
		if validateForMinEatingSpeed(piles, h, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}

func validateForMinEatingSpeed(piles []int, h int, speed int) bool {
	var count int
	for _, p := range piles {
		count += int(math.Ceil(float64(p) / float64(speed)))
	}
	return count <= h
}
