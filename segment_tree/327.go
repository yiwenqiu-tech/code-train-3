package tree_array

import (
	"sort"
)

func countRangeSum(nums []int, lower int, upper int) int {
	n := len(nums)
	var sums = make([]int, n+1)
	for i := 1; i <= n; i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}

	// 离散化
	var values []int
	for i := 0; i <= n; i++ {
		values = append(values, sums[i])
		values = append(values, sums[i]-lower)
		values = append(values, sums[i]-upper)
	}

	// 排序+去重
	sort.Ints(values)
	m := 0
	for i := 0; i < len(values); i++ {
		if i == 0 || values[i] != values[i-1] {
			values[m] = values[i]
			m++
		}
	}
	values = values[:m]

	segmentTree := NewSegmentTree327(m)
	segmentTree.Add(getIndexByValue(values, sums[0]), 1)

	var ans int
	for i := 1; i <= n; i++ {
		left := sums[i] - upper
		right := sums[i] - lower
		ans += segmentTree.Query(getIndexByValue(values, left), getIndexByValue(values, right))

		segmentTree.Add(getIndexByValue(values, sums[i]), 1)

		// O(N)，需要优化，否则超时
		//for j := 0; j < i; j++ {
		//	if sums[i] - sums[j] >= lower && sums[i] - sums[j] <= upper {
		//		ans++
		//	}
		//}
	}
	return ans
}

func getIndexByValue(values []int, val int) int {
	l := 0
	r := len(values) - 1
	for l <= r {
		mid := (r + l) / 2
		if values[mid] == val {
			return mid
		}
		if values[mid] < val {
			l = mid + 1
		}
		if values[mid] > val {
			r = mid - 1
		}
	}
	return -1
}

type Node327 struct {
	l    int
	r    int
	data int
}

type SegmentTree327 struct {
	NodeList []Node327
}

func NewSegmentTree327(m int) *SegmentTree327 {
	s := &SegmentTree327{}
	s.NodeList = make([]Node327, 4*m)
	s.buildTree(1, 0, m)
	return s
}

func (s *SegmentTree327) Add(index int, value int) {
	s.add(1, index, value)
}

func (s *SegmentTree327) Query(left int, right int) int {
	return s.query(1, left, right)
}

func (s *SegmentTree327) buildTree(p int, l int, r int) {
	s.NodeList[p].l = l
	s.NodeList[p].r = r
	if l == r {
		s.NodeList[p].data = 0
		return
	}
	mid := (r + l) / 2
	s.buildTree(p*2, l, mid)
	s.buildTree(p*2+1, mid+1, r)
	s.NodeList[p].data = s.NodeList[p*2].data + s.NodeList[p*2+1].data
}

func (s *SegmentTree327) add(p int, index int, value int) {
	l := s.NodeList[p].l
	r := s.NodeList[p].r
	if l == r {
		s.NodeList[p].data += value
		return
	}
	mid := (l + r) / 2
	if index <= mid {
		s.add(p*2, index, value)
	} else {
		s.add(p*2+1, index, value)
	}
	s.NodeList[p].data = s.NodeList[p*2].data + s.NodeList[p*2+1].data
}

func (s *SegmentTree327) query(p int, left int, right int) int {
	l := s.NodeList[p].l
	r := s.NodeList[p].r

	if left <= l && right >= r { // 完全覆盖，直接返回
		return s.NodeList[p].data
	}
	mid := (l + r) / 2
	var ans int
	if left <= mid {
		ans += s.query(p*2, left, right)
	}

	if right > mid {
		ans += s.query(p*2+1, left, right)
	}
	return ans

}
