package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var numCount int
	fmt.Scan(&numCount)

	var nums = make([]int, numCount)
	for i := 0; i < numCount; i++ {
		fmt.Scan(&nums[i])
	}

	sort.Ints(nums)

	mid := nums[len(nums)/2]

	var ans int
	for i := 0; i < len(nums); i++ {
		ans += int(math.Abs(float64(nums[i] - mid)))
	}

	fmt.Println(ans)
}
