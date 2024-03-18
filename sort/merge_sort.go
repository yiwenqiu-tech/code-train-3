package sort

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	nums1 := mergeSort(nums[0 : len(nums)/2])
	nums2 := mergeSort(nums[len(nums)/2:])
	return merge(nums1, nums2)
}

func merge(nums1 []int, nums2 []int) []int {
	if len(nums1) < 1 { // nums1为空，直接返回nums2
		return nums2
	} else if len(nums2) < 1 { // nums2为空，直接返回nums1
		return nums1
	} else { // 均不为空
		var res []int
		nums1Index := 0                                          // nums1的指针
		nums2Index := 0                                          // nums2的指针
		for nums1Index < len(nums1) && nums2Index < len(nums2) { // 均合法时，依次比较
			if nums1[nums1Index] <= nums2[nums2Index] {
				res = append(res, nums1[nums1Index])
				nums1Index++
			} else {
				res = append(res, nums2[nums2Index])
				nums2Index++
			}
		}
		if nums1Index >= len(nums1) { // nums1已经遍历完了，直接将num2放到res后面
			for i := nums2Index; i < len(nums2); i++ {
				res = append(res, nums2[i])
			}
		}
		if nums2Index >= len(nums2) { // nums2已经遍历完了，直接将num1放到res后面
			for i := nums1Index; i < len(nums1); i++ {
				res = append(res, nums1[i])
			}
		}
		return res
	}
}
