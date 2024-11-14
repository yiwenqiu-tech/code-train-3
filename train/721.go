package train

import "sort"

type UnionSet721 struct {
	fa []int
}

func (u *UnionSet721) FindFa(x int) int {
	if x == u.fa[x] {
		return x
	}
	fafa := u.FindFa(u.fa[x])
	u.fa[x] = fafa
	return fafa
}

func (u *UnionSet721) UnionSet(x, y int) {
	xFa := u.FindFa(x)
	yFa := u.FindFa(y)
	if xFa == yFa {
		return
	}
	u.fa[yFa] = xFa
}

func accountsMerge(accounts [][]string) [][]string {
	var emailMap = make(map[string]int)
	var indexNameMap = make(map[int]string)
	var us UnionSet721
	for _, account := range accounts {
		name := account[0]
		usLen := len(us.fa)
		indexs := emailInMap(account[1:], emailMap)
		if len(indexs) == 0 { // 没有添加过
			indexNameMap[usLen] = name
			us.fa = append(us.fa, usLen)
			for _, a := range account[1:] {
				emailMap[a] = usLen
			}
		} else { // 已经添加过
			fa := us.FindFa(indexs[0])
			if len(indexs) > 1 {
				for _, i := range indexs[1:] {
					us.UnionSet(fa, i)
				}
			}
			for _, a := range account[1:] {
				emailMap[a] = fa
			}
		}
	}
	var indexEmailMap = make(map[int][]string)
	for email, index := range emailMap {
		fa := us.FindFa(index) // 找到index的父节点
		if _, exist := indexEmailMap[fa]; exist {
			indexEmailMap[fa] = append(indexEmailMap[fa], email)
		} else {
			indexEmailMap[fa] = []string{email}
		}
	}
	var res [][]string
	for index, email := range indexEmailMap {
		var singleRes []string
		name := indexNameMap[index]
		singleRes = append(singleRes, name)
		sort.Strings(email)
		singleRes = append(singleRes, email...)
		res = append(res, singleRes)
	}
	return res
}

func emailInMap(emails []string, emailMap map[string]int) []int {
	var indexs []int
	for _, e := range emails {
		if index, exist := emailMap[e]; exist {
			indexs = append(indexs, index)
		}
	}
	return indexs
}
