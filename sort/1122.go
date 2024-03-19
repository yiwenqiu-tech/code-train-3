package sort

import "sort"

type intSlice struct {
	inputArr []int
	referArr map[int]int
}

func (s *intSlice) Len() int {
	return len(s.inputArr)
}

func (s *intSlice) Less(i, j int) bool {
	irIndex, iexist := s.referArr[s.inputArr[i]]
	jrIndex, jexist := s.referArr[s.inputArr[j]]
	if iexist && jexist {
		return irIndex < jrIndex
	} else if iexist {
		return true
	} else if jexist {
		return false
	} else {
		return s.inputArr[i] < s.inputArr[j]
	}
}

func (s *intSlice) Swap(i, j int) {
	s.inputArr[i], s.inputArr[j] = s.inputArr[j], s.inputArr[i]
}

func relativeSortArray(arr1 []int, arr2 []int) []int {
	referMap := make(map[int]int)
	for index, arr := range arr2 {
		referMap[arr] = index
	}
	s := &intSlice{
		inputArr: arr1,
		referArr: referMap,
	}
	sort.Sort(s)
	return s.inputArr
}
