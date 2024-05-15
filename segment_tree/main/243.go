package main

import "fmt"

func main() {
	var N int
	var M int
	fmt.Scanf("%d", &N)
	fmt.Scanf("%d", &M)

	var nums = []int{0} // 留空第0位
	for i := 0; i < N; i++ {
		var tmp int
		fmt.Scanf("%d", &tmp)
		nums = append(nums, tmp)
	}

	segmentTree := NewSegmentTree(nums)

	for i := 0; i < M; i++ {
		var operator string
		fmt.Scanf("%s", &operator)
		if operator == "Q" {
			var l int
			var r int
			fmt.Scanf("%d", &l)
			fmt.Scanf("%d", &r)
			res := segmentTree.Query(l, r)
			fmt.Println(res)
		} else {
			var l int
			var r int
			var d int
			fmt.Scanf("%d", &l)
			fmt.Scanf("%d", &r)
			fmt.Scanf("%d", &d)
			segmentTree.ChangeSeg(l, r, d)
		}
	}

}

type SegmentTreeNode struct {
	l    int
	r    int
	data int
	mark int // 延迟标记
}

type SegmentTree struct {
	tree []SegmentTreeNode
}

func NewSegmentTree(nums []int) *SegmentTree {
	s := &SegmentTree{}
	n := len(nums)
	s.tree = make([]SegmentTreeNode, 4*n)
	s.buildTree(nums, 1, 1, len(nums)-1)
	return s
}

func (s *SegmentTree) ChangeSeg(left int, right int, delta int) {
	s.changeSeg(1, left, right, delta)
}

func (s *SegmentTree) Change(index int, val int) {
	s.change(1, index, val)
}

func (s *SegmentTree) Query(left int, right int) int {
	return s.query(1, left, right)
}

func (s *SegmentTree) buildTree(nums []int, p int, l int, r int) {
	s.tree[p].l = l
	s.tree[p].r = r
	if l == r {
		s.tree[p].data = nums[l]
		return
	}
	mid := (l + r) / 2
	s.buildTree(nums, p*2, l, mid)
	s.buildTree(nums, p*2+1, mid+1, r)
	s.tree[p].data = s.tree[p*2].data + s.tree[p*2+1].data // 求sum
}

func (s *SegmentTree) changeSeg(p int, left int, right int, delta int) {
	l := s.tree[p].l
	r := s.tree[p].r
	if left <= l && r <= right { // 都包含，则返回
		s.tree[p].data += (r - l + 1) * delta
		s.tree[p].mark += delta
		return
	}
	s.spread(p)
	mid := (l + r) / 2
	if left <= mid {
		s.changeSeg(p*2, left, right, delta)
	}
	if right > mid {
		s.changeSeg(p*2+1, left, right, delta)
	}
	s.tree[p].data = s.tree[p*2].data + s.tree[p*2+1].data
}

func (s *SegmentTree) change(p int, x int, val int) {
	if s.tree[p].l == s.tree[p].r {
		s.tree[p].data = val
		return
	}
	s.spread(p)
	l := s.tree[p].l
	r := s.tree[p].r
	mid := (l + r) / 2
	if x <= mid {
		s.change(p*2, x, val)
	} else {
		s.change(p*2+1, x, val)
	}
	s.tree[p].data = s.tree[p*2].data + s.tree[p*2+1].data
}

func (s *SegmentTree) query(p int, left int, right int) int {
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

func (s *SegmentTree) spread(p int) {
	if s.tree[p].mark != 0 {
		s.tree[p*2].data += (s.tree[p*2].r - s.tree[p*2].l + 1) * s.tree[p].mark
		s.tree[p*2].mark += s.tree[p].mark
		s.tree[p*2+1].data += (s.tree[p*2+1].r - s.tree[p*2+1].l + 1) * s.tree[p].mark
		s.tree[p*2+1].mark += s.tree[p].mark
		s.tree[p].mark = 0
	}
}
