package array

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 { // 考虑num1为空，将num2复制到num1
		for i := 0; i < n; i++ {
			nums1[i] = nums2[i]
		}
		return
	}
	if n == 0 { // 考虑nums2为空，则直接返回，结果已经在nums1
		return
	}
	for m > 0 && n > 0 {
		if nums1[m-1] < nums2[n-1] { // 从最后一个数比较，若nums2大，则将其放到nums1的最后，然后将n--
			nums1[m+n-1] = nums2[n-1]
			n--
		} else {
			nums1[m+n-1] = nums1[m-1] // 从最后一个数比较，若nums1大，则将其放到nums1的最后，然后将m--
			m--
		}
	}
	for n > 0 {
		nums1[n-1] = nums2[n-1]
		n--
	}
}
