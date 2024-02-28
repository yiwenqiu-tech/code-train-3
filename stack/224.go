package stack

func calculate224(s string) int {
	s = s + " "
	if s[0] == '-' { // 若开头是-,则添加上0
		s = "0" + s
	}
	var intStack IntListStack
	var charStack CharListStack
	var curNum int
	var hasNum bool
	var lastIsOperator bool // 上一次是否为操作符
	for index := range s {
		if s[index] >= '0' && s[index] <= '9' {
			hasNum = true
			num := int(s[index] - '0')
			curNum = curNum*10 + num
			lastIsOperator = false
			continue
		} else {
			if hasNum {
				intStack.Push(curNum)
				curNum = 0 // 重置curNum
				hasNum = false
			}
			if s[index] == ' ' {
				continue
			}
			if lastIsOperator && s[index] != '(' { // 上一个是操作符且当前符号不是( （-(这种不需要插），则插0
				intStack.Push(0)
				lastIsOperator = false
			} else { // 否则标识上一个是操作符
				lastIsOperator = true
			}
			if s[index] == '(' {
				charStack.Push(s[index])
				continue
			}
			if s[index] == ')' {
				for charStack.Peek() != '(' {
					operator := charStack.Pop()
					num2 := intStack.Pop()
					num1 := intStack.Pop()
					intStack.Push(cal(operator, num1, num2))
				}
				charStack.Pop()
				lastIsOperator = false // 当当前字符为')'时，清空lastIsOperator，因为下一个字符不可能为数字的负号
				continue
			}

			if charStack.Empty() {
				charStack.Push(s[index])
			} else {
				for !charStack.Empty() && charStack.Peek() != '(' &&
					operatorRank[charStack.Peek()] >= operatorRank[s[index]] {
					operator := charStack.Pop()
					num2 := intStack.Pop()
					num1 := intStack.Pop()
					intStack.Push(cal(operator, num1, num2))
				}
				charStack.Push(s[index])
			}
		}
	}
	for !charStack.Empty() {
		operator := charStack.Pop()
		num2 := intStack.Pop()
		num1 := intStack.Pop()
		intStack.Push(cal(operator, num1, num2))
	}
	return intStack.Peek()
}
