package train

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for index, num := range nums {
		if i, exist := numMap[target-num]; exist {
			return []int{i, index}
		}
		numMap[num] = index
	}
	return nil
}
