package binary_search

var pick int

func guess(num int) int {
	if pick < num {
		return -1
	} else if pick > num {
		return 1
	} else {
		return 0
	}
}

func guessNumber(n int) int {
	left := 1
	right := n
	for left <= right {
		mid := left + (right-left)/2
		if guess(mid) == 0 {
			return mid
		}
		if guess(mid) < 0 {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
