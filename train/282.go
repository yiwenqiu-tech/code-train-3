package train

import (
	"container/list"
	"fmt"
	"sort"
)

var target282 int
var res282 []string

func addOperators(num string, target int) []string {
	target282 = target
	res282 = nil

	dfsFor282(num, 0, nil)
	sort.Strings(res282)
	fmt.Println(res282)
	return res282
}

func dfsFor282(num string, i int, cur []byte) {
	if len(num)-1 == i {
		cur = append(cur, num[i])
		if calNum(cur) == target282 {
			res282 = append(res282, string(cur))
		}
		//for i := range cur {
		//	fmt.Print(string(cur[i]))
		//}
		//fmt.Printf("\n")
		return
	}
	cur = append(cur, num[i])

	if (num[i] > '0' && num[i] <= '9') || (len(cur) >= 2 && (cur[len(cur)-2] != '-' && cur[len(cur)-2] != '+' && cur[len(cur)-2] != '*')) {
		dfsFor282(num, i+1, cur) // 什么都不加
	}

	for _, b := range []byte{'+', '-', '*'} {
		cur = append(cur, b)
		dfsFor282(num, i+1, cur)
		cur = cur[0 : len(cur)-1]
	}
	cur = cur[0 : len(cur)-1]
}

func calNum(num []byte) int { // 考虑表达式运算
	numStack := list.New()
	opStack := list.New()
	var i int
	for i < len(num) {
		if num[i] >= '0' && num[i] <= '9' {
			var value int
			for i < len(num) && num[i] >= '0' && num[i] <= '9' {
				value = value*10 + int(num[i]-'0')
				i++
			}
			numStack.PushBack(value)
		} else {
			op := num[i]
			if opStack.Len() == 0 {
				opStack.PushBack(op)
			} else {
				stackTop := opStack.Back()
				for opSort(stackTop.Value.(byte)) >= opSort(op) {
					opStack.Remove(stackTop)
					a := numStack.Back()
					aValue := a.Value.(int)
					numStack.Remove(a)
					b := numStack.Back()
					bValue := b.Value.(int)
					numStack.Remove(b)
					tmp := cal(stackTop.Value.(byte), aValue, bValue)
					numStack.PushBack(tmp)
					if opStack.Len() == 0 {
						break
					}
					stackTop = opStack.Back()
				}
				opStack.PushBack(op)
			}
			i++
		}
	}
	for opStack.Len() > 0 {
		stackTop := opStack.Back()
		opStack.Remove(stackTop)
		a := numStack.Back()
		aValue := a.Value.(int)
		numStack.Remove(a)
		b := numStack.Back()
		bValue := b.Value.(int)
		numStack.Remove(b)
		numStack.PushBack(cal(stackTop.Value.(byte), aValue, bValue))
	}
	return numStack.Back().Value.(int)
}

func cal(op byte, aValue, bValue int) int {
	if op == '*' {
		return aValue * bValue
	} else if op == '+' {
		return aValue + bValue
	} else {
		return bValue - aValue
	}
}

func opSort(op byte) int {
	if op == '*' {
		return 2
	} else {
		return 1
	}
}
