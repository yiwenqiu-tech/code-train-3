package double_pointer

func twoSum(numbers []int, target int) []int {
	if len(numbers) < 2 {
		return nil
	}
	left := 0
	right := len(numbers) - 1

	for left != right {
		if numbers[left]+numbers[right] == target {
			return []int{left + 1, right + 1}
		}
		if numbers[left]+numbers[right] < target {
			left++
		} else {
			right--
		}
	}
	return nil
}
