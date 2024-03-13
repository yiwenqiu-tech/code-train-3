package binary_search

func findPeakElement(nums []int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		lmid := left + (right-left)/2 // 取中点为左点
		rmid := lmid + 1              // 取中点+1为右点
		if nums[lmid] < nums[rmid] {
			left = lmid + 1 // 由于rmid大于lmid，所以lmid不可能是答案，因此+1
		} else {
			right = rmid - 1 // 因为题意里不会等于，同样lmid > rmid,rmid不可能是答案，因此-1
		}
	}
	return right
}
