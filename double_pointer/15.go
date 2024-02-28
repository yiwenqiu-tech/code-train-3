package double_pointer

import "sort"

func threeSum(nums []int) [][]int {
	sort.Sort(sort.IntSlice(nums)) // 先排序，方便后面去重
	var res [][]int
	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		k := 0 - nums[i]
		if len(nums)-(i+1) < 2 {
			return res
		}
		if twoRes := twoSumForTarget(nums, i+1, k); twoRes != nil {
			for _, r := range twoRes {
				r = append(r, nums[i])
				res = append(res, r)
			}
		}
	}
	return res
}

func twoSumForTarget(nums []int, start int, target int) [][]int {
	var res [][]int
	end := len(nums) - 1
	tmpStart := start
	// 已排序的，通过双指针来查找
	for tmpStart != end {
		if tmpStart != start && nums[tmpStart] == nums[tmpStart-1] {
			tmpStart++
			continue
		}
		if nums[tmpStart]+nums[end] == target {
			res = append(res, []int{nums[tmpStart], nums[end]})
			tmpStart++
			continue
		}
		if nums[tmpStart]+nums[end] > target {
			end--
		} else {
			tmpStart++
		}
	}
	return res
}
