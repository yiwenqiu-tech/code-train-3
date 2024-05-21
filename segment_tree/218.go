package tree_array

import (
	"sort"
)

func getSkyline(buildings [][]int) [][]int {
	var buildPos []int
	for _, building := range buildings {
		buildPos = append(buildPos, building[0])
		buildPos = append(buildPos, building[1])
	}
	sort.Ints(buildPos)

	var buildPosIndexMap = make(map[int]int)
	var indexBuildPosMap = make(map[int]int)
	for index, pos := range buildPos {
		buildPosIndexMap[pos] = index
		indexBuildPosMap[index] = pos
	}
	length := len(buildPos)
	segment := NewSegmentTree218(length)
	for _, building := range buildings {
		segment.ChangeSeg(buildPosIndexMap[building[0]], buildPosIndexMap[building[1]]-1, building[2])
	}

	var latest = 0
	var ans [][]int
	for i := 0; i < length; i++ {
		temp := segment.Query(i, i)
		if temp == latest {
			continue
		}
		latest = temp
		ans = append(ans, []int{indexBuildPosMap[i], temp})
	}
	return ans
}

type SegmentTreeNode218 struct {
	l    int
	r    int
	data int
	mark int // 延迟标记
}

type SegmentTree218 struct {
	tree []SegmentTreeNode218
}

func NewSegmentTree218(length int) *SegmentTree218 {
	s := &SegmentTree218{}
	n := length
	s.tree = make([]SegmentTreeNode218, 4*n)
	s.buildTree(1, 0, length)
	return s
}

func (s *SegmentTree218) ChangeSeg(left int, right int, data int) {
	s.changeSeg(1, left, right, data)
}

func (s *SegmentTree218) Query(left int, right int) int {
	return s.query(1, left, right)
}

func (s *SegmentTree218) buildTree(p int, l int, r int) {
	s.tree[p].l = l
	s.tree[p].r = r
	if l == r {
		s.tree[p].data = 0
		return
	}
	mid := (l + r) / 2
	s.buildTree(p*2, l, mid)
	s.buildTree(p*2+1, mid+1, r)
}

func (s *SegmentTree218) changeSeg(p int, left int, right int, data int) {
	l := s.tree[p].l
	r := s.tree[p].r
	if left <= l && r <= right { // 都包含，则返回
		if data > s.tree[p].data {
			s.tree[p].data = data
		}
		if data > s.tree[p].mark { // 大于现有的值，则修改mark
			s.tree[p].mark = data
		}
		return
	}
	s.spread(p)
	mid := (l + r) / 2
	if left <= mid {
		s.changeSeg(p*2, left, right, data)
	}
	if right > mid {
		s.changeSeg(p*2+1, left, right, data)
	}
	s.tree[p].data = max(s.tree[p*2].data, s.tree[p*2+1].data)
}

func (s *SegmentTree218) query(p int, left int, right int) int {
	l := s.tree[p].l
	r := s.tree[p].r
	if left <= l && r <= right { // 都包含，则返回
		return s.tree[p].data
	}
	s.spread(p)
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

func (s *SegmentTree218) spread(p int) {
	if s.tree[p].mark != 0 {
		s.tree[p*2].data = max(s.tree[p*2].data, s.tree[p].mark)
		if s.tree[p*2].mark < s.tree[p].mark {
			s.tree[p*2].mark = s.tree[p].mark
		}
		s.tree[p*2+1].data = max(s.tree[p*2+1].data, s.tree[p].mark)
		if s.tree[p*2+1].mark < s.tree[p].mark {
			s.tree[p*2+1].mark = s.tree[p].mark
		}
		s.tree[p].mark = 0
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
