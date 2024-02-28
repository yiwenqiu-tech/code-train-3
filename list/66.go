package list

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			break
		} else {
			digits[i] = 0
		}
	}

	if digits[0] != 0 {
		return digits
	}

	var res = make([]int, len(digits)+1)
	res[0] = 1
	for i := 0; i < len(digits)-1; i++ {
		res[i+1] = digits[i]
	}
	return res
}
