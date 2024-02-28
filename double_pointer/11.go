package double_pointer

import "math"

func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}

	left := 0
	right := len(height) - 1

	var res = math.MinInt32
	for left != right {
		tmp := min(height[left], height[right]) * (right - left)
		if tmp > res {
			res = tmp
		}
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}

	return res
}

func min(left, right int) int {
	if left < right {
		return left
	}

	return right
}
