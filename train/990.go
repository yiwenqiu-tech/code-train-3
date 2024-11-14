package train

import "strings"

type UnionSet struct {
	fa []int
}

func (u *UnionSet) FindFa(n int) int {
	if n == u.fa[n] {
		return n
	}
	fafa := u.FindFa(u.fa[n])
	u.fa[n] = fafa
	return fafa
}

func (u *UnionSet) UnionSet(x, y int) {
	xFa := u.FindFa(x)
	yFa := u.FindFa(y)
	if xFa == yFa {
		return
	}
	u.fa[yFa] = xFa
}

func equationsPossible(equations []string) bool {
	equMap := make(map[string]int)
	var us UnionSet
	var noEqualSts [][]string
	for _, str := range equations {
		strs, operation := splitEquation(str)
		for _, s := range strs { // 添加到并查集里
			if _, exist := equMap[s]; !exist {
				index := len(us.fa)
				us.fa = append(us.fa, index)
				equMap[s] = index
			}
		}
		if operation == "==" { // 先处理等于的，将所有等于的放到一个并查集里
			us.UnionSet(equMap[strs[0]], equMap[strs[1]])
		} else { // 不等于的后面处理
			noEqualSts = append(noEqualSts, strs)
		}
	}
	for _, strs := range noEqualSts { // 处理不等于的
		xStrFa := us.FindFa(equMap[strs[0]])
		yStrFa := us.FindFa(equMap[strs[1]])
		if xStrFa == yStrFa {
			return false
		}
	}

	return true
}

func splitEquation(equ string) ([]string, string) {
	res := strings.Split(equ, "==")
	operation := "=="
	if len(res) != 2 {
		res = strings.Split(equ, "!=")
		operation = "!="
	}
	return res, operation
}
