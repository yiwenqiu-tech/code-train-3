package stack

var operatorRank = map[byte]int{
	'*': 2,
	'/': 2,
	'+': 1,
	'-': 1,
}

func calculate(s string) int {
	s = s + " "
	var intStack IntListStack
	var charStack CharListStack
	var curNum int
	var hasNum bool
	for index := range s {
		if s[index] >= '0' && s[index] <= '9' {
			hasNum = true
			num := int(s[index] - '0')
			curNum = curNum*10 + num
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

			if charStack.Empty() {
				charStack.Push(s[index])
			} else {
				for !charStack.Empty() && operatorRank[charStack.Peek()] >= operatorRank[s[index]] {
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

func cal(operator byte, num1, num2 int) int {
	switch operator {
	case '*':
		return num1 * num2
	case '/':
		return num1 / num2
	case '+':
		return num1 + num2
	case '-':
		return num1 - num2
	}
	return 0
}
