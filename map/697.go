package _map

import "math"

func findShortestSubArray(nums []int) int {
	var numsCount = make(map[int]int)
	var numsFirst = make(map[int]int)
	var numsLast = make(map[int]int)
	var maxCount int
	for index, num := range nums {
		if _, exist := numsFirst[num]; !exist {
			numsFirst[num] = index
		}
		numsLast[num] = index
		numsCount[num]++

		if numsCount[num] > maxCount {
			maxCount = numsCount[num]
		}
	}
	var minLen = math.MaxInt32

	for num, count := range numsCount {
		if count == maxCount {
			arrLen := numsLast[num] - numsFirst[num] + 1
			if arrLen < minLen {
				minLen = arrLen
			}
		}
	}
	return minLen
}
