package tree_array

type TreeArray struct { // 树状数组
	a []int
	c []int
}

func NewTreeArray(input []int) *TreeArray {
	n := len(input)
	t := &TreeArray{}
	t.c = make([]int, n+1) // 第一位留空
	t.a = make([]int, n+1) // 第一位留空
	for i := 1; i <= n; i++ {
		t.add(i, input[i-1])
	}

	return t
}

func (t *TreeArray) queryPrefixSum(x int) int {
	ans := 0
	for ; x > 0; x -= t.lowbit(x) {
		ans += t.c[x]
	}
	return ans
}

func (t *TreeArray) lowbit(x int) int {
	return x & -x
}

func (t *TreeArray) add(index int, value int) {
	i := index
	for ; i < len(t.c); i += t.lowbit(i) { // 更新树状数组
		t.c[i] += value
	}
	t.a[index] += value // 更新原数组
}

type NumArray struct {
	t *TreeArray
}

func Constructor(nums []int) NumArray {
	return NumArray{
		t: NewTreeArray(nums),
	}
}

func (this *NumArray) Update(index int, val int) {
	index++
	this.t.add(index, val-this.t.a[index])
}

func (this *NumArray) SumRange(left int, right int) int {
	left++
	right++
	return this.t.queryPrefixSum(right) - this.t.queryPrefixSum(left-1)
}
