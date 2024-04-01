package sort

import "fmt"

var rangeLower int
var rangeUpper int
var rangeSum int

func countRangeSum(nums []int, lower int, upper int) int {
	rangeSum = 0
	rangeLower = lower
	rangeUpper = upper
	sums := make([]int64, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		sums[i+1] = sums[i] + int64(nums[i])
	} // 前缀和
	fmt.Println(sums)
	mergeSortForCountRangeSum(sums, 0, len(sums)-1)
	return rangeSum
}

func mergeSortForCountRangeSum(sums []int64, left, right int) {
	if left >= right {
		return
	}
	mid := left + (right-left)/2
	mergeSortForCountRangeSum(sums, left, mid)
	mergeSortForCountRangeSum(sums, mid+1, right)
	calRangeCount(sums, left, mid, right)
	mergeForCountRangeSum(sums, left, mid, right)
}

func calRangeCount(nums []int64, left, mid, right int) {
	rl := 0
	rr := 0
	for i := left; i <= mid; i++ {
		for rl+mid+1 <= right && nums[rl+mid+1]-nums[i] < int64(rangeLower) {
			rl++
		} // 第一个满足 >= rangeLower的
		if rr == 0 { // 初始化时记录为rl
			rr = rl
		}
		for rr+mid+1 <= right && nums[rr+mid+1]-nums[i] <= int64(rangeUpper) {
			rr++
		} // 第一个不满足 <=rangeUpper 的
		if rl >= rr {
			continue
		}
		rangeSum += (rr - 1) - rl + 1
		fmt.Println(rangeSum)
	}
}

func mergeForCountRangeSum(nums []int64, left, mid, right int) {
	var index1 = left
	var index2 = mid + 1
	var tmp []int64
	for index1 <= mid && index2 <= right {
		if nums[index1] < nums[index2] {
			tmp = append(tmp, nums[index1])
			index1++
		} else {
			tmp = append(tmp, nums[index2])
			index2++
		}
	}
	for index1 <= mid {
		tmp = append(tmp, nums[index1])
		index1++
	}
	for index2 <= right {
		tmp = append(tmp, nums[index2])
		index2++
	}
	for i := left; i <= right; i++ {
		nums[i] = tmp[i-left]
	}
}
