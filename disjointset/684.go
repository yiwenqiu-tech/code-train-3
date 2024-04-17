package disjointset

type ds684 []int

func (d *ds684) find(x int) int {
	if (*d)[x] == x {
		return x
	}
	fa := d.find((*d)[x])
	(*d)[x] = fa
	return fa
}

func findRedundantConnection(edges [][]int) []int {
	fa := make(ds684, len(edges)+1) // make point
	for i := 1; i <= len(edges); i++ {
		fa[i] = i
	}
	for i := range edges {
		root0 := fa.find(edges[i][0])
		root1 := fa.find(edges[i][1])
		if root0 == root1 {
			return edges[i]
		}
		fa[root0] = root1 // 合并并查集
	}
	return nil
}
