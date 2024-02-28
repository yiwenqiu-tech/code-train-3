package _map

func twoSum(nums []int, target int) []int {
	var numMap = make(map[int]int)
	for index, num := range nums {
		if value, exist := numMap[target-num]; exist {
			return []int{value, index}
		}
		numMap[num] = index
	}
	return nil
}
