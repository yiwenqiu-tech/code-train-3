package sort

import (
	"math"
	"sort"
)

type slices [][]int

func (s *slices) Len() int {
	return len(*s)
}

func (s *slices) Less(i, j int) bool {
	if (*s)[i][0] < (*s)[j][0] {
		return true
	} else if (*s)[i][0] == (*s)[j][0] {
		return (*s)[i][1] < (*s)[j][1]
	} else {
		return false
	}
}

func (s *slices) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

func merge56(intervals [][]int) [][]int {
	input := slices(intervals)
	sort.Sort(&input)

	var ans [][]int
	var begin = -1
	var end = -1
	for i := 0; i < input.Len(); i++ {
		if begin == -1 {
			begin = input[i][0]
			end = input[i][1]
			continue
		}
		if input[i][0] <= end {
			end = int(math.Max(float64(input[i][1]), float64(end))) // 更新end
		} else {
			// 如果大于上一段的末尾，则新开一段，将之前那段放到答案里，并更新begin和end
			ans = append(ans, []int{begin, end})
			begin = input[i][0]
			end = input[i][1]
		}
	}
	ans = append(ans, []int{begin, end})
	return ans
}

type diffSlice [][]int

func (s *diffSlice) Len() int {
	return len(*s)
}

func (s *diffSlice) Less(i, j int) bool {
	if (*s)[i][0] < (*s)[j][0] {
		return true
	} else if (*s)[i][0] == (*s)[j][0] {
		return (*s)[i][1] < (*s)[j][1]
	} else {
		return false
	}
}

func (s *diffSlice) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

func merge562(intervals [][]int) [][]int {
	var diffArr [][]int
	for _, i := range intervals {
		diffArr = append(diffArr, []int{i[0], 1})
		diffArr = append(diffArr, []int{i[1] + 1, -1}) // 差分，在后面的值上+1
	}
	dSlice := diffSlice(diffArr)
	sort.Sort(&dSlice)
	var ans [][]int
	var begin = -1
	var cur = 0
	for _, s := range dSlice {
		if begin == -1 && s[1] == 1 {
			begin = s[0]
			cur++
			continue
		}
		if s[1] == 1 {
			cur++
		} else {
			cur--
		}
		if cur == 0 {
			ans = append(ans, []int{begin, s[0] - 1})
			begin = -1
		}
	}
	return ans
}
