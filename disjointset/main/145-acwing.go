package main

import (
	"fmt"
	"io"
	"sort"
)

type product struct {
	income int
	expire int
}

type productSlice []product

func (p *productSlice) Len() int {
	return len(*p)
}

func (p *productSlice) Less(i, j int) bool {
	return (*p)[i].income > (*p)[j].income
}

func (p *productSlice) Swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

type disjointSet [10001]int

func (ds *disjointSet) find(x int) int {
	if x == ds[x] {
		return x
	}
	fa := ds.find(ds[x])
	ds[x] = fa
	return fa
}

func main() {
	for {
		var n = -1 // 数字数量
		_, err := fmt.Scanf("%d", &n)
		if err == io.EOF { // 没有输入了直接break
			break
		}
		if n == -1 {
			continue
		}
		products := make([]product, n)
		for i := 0; i < n; i++ {
			fmt.Scanf("%d %d", &products[i].income, &products[i].expire)
		}
		p := productSlice(products)
		sort.Sort(&p)
		var ds disjointSet // makeSet
		for i := range ds {
			ds[i] = i
		}
		var ans int
		for i := range p {
			validDay := ds.find(p[i].expire) // 寻找空闲日
			if validDay != 0 {               // 存在空闲日
				ans += p[i].income
				ds[validDay] = validDay - 1 // 直接指向上一个并查集
			}
		}
		fmt.Println(ans)
	}

}
