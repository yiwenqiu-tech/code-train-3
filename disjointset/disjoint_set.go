package disjointset

type DisjointSet struct {
	Fa []int
}

func MakeSet(n int) *DisjointSet {
	initFa := make([]int, n)
	for i := range initFa {
		initFa[i] = i // 初始化时父节点指向自己
	}
	return &DisjointSet{
		Fa: initFa,
	}
}

// Find 返回其根结点，查找过程中进行路径压缩
func (d *DisjointSet) Find(n int) int {
	if n == d.Fa[n] {
		return n
	}
	fafa := d.Find(d.Fa[n]) // 找出当前父亲的父亲
	d.Fa[n] = fafa          //逐层回溯到n，最后每层都赋值为fafa，达到路径压缩的效果
	return fafa
}

// UnionSet 合并
func (d *DisjointSet) UnionSet(x int, y int) {
	xRoot := d.Find(x)
	yRoot := d.Find(y)
	if xRoot == yRoot { // x,y的根节点
		return
	}
	d.Fa[yRoot] = xRoot // 将y的fa设置为xRoot
}
