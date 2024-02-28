package recur

import (
	"sort"
)

var permuteUniqueRes [][]int
var permuteUniqueChosen []int
var permuteUniqueChosenBool []bool

func permuteUnique(nums []int) [][]int {
	permuteUniqueRes = nil
	permuteUniqueChosenBool = make([]bool, len(nums))
	sort.Sort(sort.IntSlice(nums)) // 先排序
	recurForPermuteUnique2(nums)
	return permuteUniqueRes
}

func recurForPermuteUnique2(nums []int) {
	if len(permuteUniqueChosen) == len(nums) {
		var tmpRes = make([]int, len(permuteUniqueChosen))
		copy(tmpRes, permuteUniqueChosen)
		permuteUniqueRes = append(permuteUniqueRes, tmpRes)
		return
	}
	for i := 0; i < len(nums); i++ {
		if permuteUniqueChosenBool[i] {
			continue
		}
		// 剪枝技巧，若在同一层中，当前所选的数与前一个数一致，且前一个数还没被选过，则可以剪掉。【注意：需要先对nums排序】
		if i > 0 && (nums[i] == nums[i-1]) && !permuteUniqueChosenBool[i-1] {
			continue
		}
		permuteUniqueChosen = append(permuteUniqueChosen, nums[i])
		permuteUniqueChosenBool[i] = true
		recurForPermuteUnique2(nums)
		permuteUniqueChosen = permuteUniqueChosen[:len(permuteUniqueChosen)-1]
		permuteUniqueChosenBool[i] = false
	}
}

//func recurForPermuteUnique(nums []int) {
//	if len(permuteUniqueChosen) == len(nums) {
//		// 暴力做法，把所有结果遍历出来后，再进行去重。
//		if _, exist := permuteUniqueChosenMap[genMapKey(permuteUniqueChosen)]; !exist {
//			var tmpRes = make([]int, len(permuteUniqueChosen))
//			copy(tmpRes, permuteUniqueChosen)
//			permuteUniqueRes = append(permuteUniqueRes, tmpRes)
//			permuteUniqueChosenMap[genMapKey(permuteUniqueChosen)] = struct{}{}
//		}
//		return
//	}
//	for i := 0; i < len(nums); i++ {
//		if !permuteUniqueChosenBool[i] {
//			permuteUniqueChosen = append(permuteUniqueChosen, nums[i])
//			permuteUniqueChosenBool[i] = true
//			recurForPermuteUnique(nums)
//			permuteUniqueChosen = permuteUniqueChosen[:len(permuteUniqueChosen)-1]
//			permuteUniqueChosenBool[i] = false
//		}
//	}
//}
//
//func genMapKey(nums []int) string {
//	var key string
//	for i := 0; i < len(nums); i++ {
//		key += strconv.Itoa(nums[i]) + "_"
//	}
//	return key
//}
