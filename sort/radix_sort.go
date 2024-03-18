package sort

import "math"

func RadixSort(nums []int) []int {
	maxInt := math.MinInt32
	minInt := math.MaxInt32
	for _, num := range nums {
		if num > maxInt {
			maxInt = num
		}
		if num < minInt {
			minInt = num
		}
	}

	// 考虑数组存在负数时，先都调整为正数，注意maxInt也要调整
	if minInt < 0 {
		for index := range nums {
			nums[index] -= minInt
		}
		maxInt -= minInt
	}

	maxDigit := 0
	for maxInt > 0 {
		maxDigit++
		maxInt = maxInt / 10
	}

	// 外层循环maxDigit
	for i := 0; i < maxDigit; i++ {
		var count = make([][]int, 10)
		// 里层根据所在位数的数值决定放到那个桶里
		for _, num := range nums {
			count[(num/int(math.Pow(float64(10), float64(i))))%10] = append(count[(num/int(math.Pow(float64(10), float64(i))))%10], num)
		}
		// 依次放回到nums里，参与到下一轮排序
		index := 0
		for _, c := range count {
			for _, n := range c {
				nums[index] = n
				index++
			}
		}
	}
	// 如果数组里存在负数，最后排好序后再调整回去原值
	if minInt < 0 {
		for index := range nums {
			nums[index] += minInt
		}
	}
	return nums
}
