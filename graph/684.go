package graph

import "container/list"

var to [][]int

var maxPoint int

var visit []bool

func findRedundantConnectionA(edges [][]int) []int {
	for _, edge := range edges {
		a := edge[0]
		b := edge[1]
		if a > maxPoint {
			maxPoint = a
		}
		if b > maxPoint {
			maxPoint = b
		}
	}
	to = make([][]int, maxPoint+1) // 1-n有效，0位置无效
	for _, edge := range edges {   // 每次加一条边，边加边判断是否成环了，当遇到成环时，返回对应的边。
		visit = make([]bool, maxPoint+1)
		a := edge[0]
		b := edge[1]
		to[a] = append(to[a], b)
		to[b] = append(to[b], a)
		if dfsForFindRedundantConnection(a, 0) != nil {
			return edge
		}
	}
	return nil
}

func findRedundantConnection(edges [][]int) []int {
	for _, edge := range edges {
		a := edge[0]
		b := edge[1]
		if a > maxPoint {
			maxPoint = a
		}
		if b > maxPoint {
			maxPoint = b
		}
	}
	to = make([][]int, maxPoint+1) // 1-n有效，0位置无效
	in := make([]int, maxPoint+1)  // 入度数组，1-n有效，0位置无效
	for _, edge := range edges {   // 构造出边数组和入度数组
		visit = make([]bool, maxPoint+1)
		a := edge[0]
		b := edge[1]
		to[a] = append(to[a], b)
		to[b] = append(to[b], a)
		in[a]++
		in[b]++
	}
	queue := list.List{}
	for index, i := range in {
		if index == 0 {
			continue
		}
		if i == 1 { // 无向连通图，1表示只有1个节点与之相连（不会在环里），找出这类节点入队
			queue.PushBack(index)
		}
	}
	for queue.Len() != 0 {
		head := queue.Front()
		headVal := head.Value.(int)
		queue.Remove(head)

		toOut := to[headVal]
		for _, out := range toOut {
			in[out]--
			if in[out] == 1 {
				queue.PushBack(out)
			}
		}
	}
	for i := len(edges) - 1; i >= 0; i-- {
		a := edges[i][0]
		b := edges[i][1]
		if in[a] > 1 && in[b] > 1 {
			return edges[i]
		}
	}
	return nil
}

func dfsForFindRedundantConnection(start, parent int) []int {
	startTo := to[start]
	visit[start] = true
	for _, i := range startTo {
		if i == parent {
			continue
		}
		if visit[i] {
			return []int{start, i}
		}
		res := dfsForFindRedundantConnection(i, start)
		if res != nil {
			return res
		}
	}
	return nil
}
