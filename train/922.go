package train

func sortArrayByParityII(nums []int) []int {
	// 方法1：额外空间解决了
	//var res = make([]int, len(nums))
	//var even = 0
	//var odd = 1
	//for _, num := range nums {
	//	if num&1 == 0 {
	//		res[even] = num
	//		even += 2
	//	} else {
	//		res[odd] = num
	//		odd += 2
	//	}
	//}
	//return res

	// 方法2：双指针
	evenPointer := 0
	oddPointer := 1
	for {
		for evenPointer < len(nums) && nums[evenPointer]&1 == 0 {
			evenPointer += 2
		}
		if evenPointer >= len(nums) {
			break
		}
		for oddPointer < len(nums) && nums[oddPointer]&1 == 1 {
			oddPointer += 2
		}
		if oddPointer >= len(nums) {
			break
		}
		nums[oddPointer], nums[evenPointer] = nums[evenPointer], nums[oddPointer]
		oddPointer += 2
		evenPointer += 2
	}
	return nums
}
