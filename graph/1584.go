package graph

import (
	"math"
	"sort"
)

type edge struct {
	x    int
	y    int
	cost int
}

type disjointSet []int

func (d *disjointSet) find(x int) int {
	if (*d)[x] == x {
		return x
	}
	fa := d.find((*d)[x])
	(*d)[x] = fa
	return fa
}

func minCostConnectPoints(points [][]int) int {
	var edges []edge
	var fa = make(disjointSet, len(points))
	for i := 0; i < len(points); i++ {
		fa[i] = i
		for j := i; j < len(points); j++ {
			edges = append(edges, edge{
				x:    i,
				y:    j,
				cost: int(math.Abs(float64(points[i][0])-float64(points[j][0])) + math.Abs(float64(points[i][1])-float64(points[j][1]))),
			})
		}
	}
	// 从小到大排序
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].cost < edges[j].cost
	})

	var ans int
	for _, e := range edges {
		if fa.find(e.x) == fa.find(e.y) { // 在一个并查集，成环了
			continue
		}
		// 否则添加到结果里
		ans += e.cost
		fa[fa.find(e.x)] = fa.find(e.y) // 更新并查集
	}
	return ans
}
