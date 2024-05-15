package tree_array

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
	s.buildTree(nums, 1, 0, len(nums)-1)
	return s
}

func (s *SegmentTree) Change(index int, val int) {
	s.change(1, index, val)
}

func (s *SegmentTree) ChangeSeg(left int, right int, delta int) {
	s.changeSeg(1, left, right, delta)
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

type NumArray struct {
	s *SegmentTree
}

func Constructor(nums []int) NumArray {
	return NumArray{
		s: NewSegmentTree(nums),
	}
}

func (this *NumArray) Update(index int, val int) {
	this.s.Change(index, val)
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.s.Query(left, right)
}
