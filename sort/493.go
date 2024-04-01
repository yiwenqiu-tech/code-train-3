package sort

var pairCount int

func reversePairs(nums []int) int {
	pairCount = 0
	mergeSortForReversePairs(nums, 0, len(nums)-1)
	return pairCount
}

func mergeSortForReversePairs(nums []int, left, right int) {
	if left >= right {
		return
	}
	mid := left + (right-left)/2
	mergeSortForReversePairs(nums, left, mid)
	mergeSortForReversePairs(nums, mid+1, right)
	calPairCount(nums, left, mid, right)
	mergeForReversePairs(nums, left, mid, right)
}

func calPairCount(nums []int, left, mid, right int) {
	j := 0
	for i := left; i <= mid; i++ {
		for j+mid+1 <= right && int64(nums[i]) > 2*(int64(nums[j+mid+1])) {
			j++
		}
		pairCount += j
	}
}

func mergeForReversePairs(nums []int, left, mid, right int) {
	var index1 = left
	var index2 = mid + 1
	var tmp []int
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
