package train

import (
	"strconv"
	"strings"
)

func decodeString(s string) string {
	var index = 0
	stack := NewStack()
	for index < len(s) {
		if isDigit(s[index]) {
			nextIndex, digitStr := getDigit(s, index)
			stack.Push(digitStr)
			index = nextIndex
			continue
		} else if s[index] == '[' || isLowerCaseLetter(s[index]) {
			stack.Push(string(s[index]))
			index++
			continue
		} else {
			index++
			str := getStrFromStack(stack)
			repeat := stack.Pop().(string)
			repeatCount, _ := strconv.Atoi(repeat)
			repeatStr := strings.Repeat(str, repeatCount)
			stack.Push(repeatStr)
		}
	}
	var ansStack = NewStack()
	var ans string
	for stack.Len() > 0 {
		ansStack.Push(stack.Pop())
	}
	for ansStack.Len() > 0 {
		ans += ansStack.Pop().(string)
	}
	return ans
}

func getStrFromStack(s *Stack) string {
	var res string
	var newS = NewStack()
	for s.Len() > 0 {
		tmp := s.Pop().(string)
		if tmp == "[" { // 遇到左括号则退出
			break
		}
		newS.Push(tmp)
	}
	for newS.Len() > 0 {
		res += newS.Pop().(string)
	}
	return res
}

func getDigit(s string, index int) (int, string) {
	tmp := 0
	for isDigit(s[index]) {
		tmp = tmp*10 + int(s[index]-'0')
		index++
	}
	return index, strconv.Itoa(tmp)
}

func isDigit(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

func isLowerCaseLetter(c byte) bool {
	if c >= 'a' && c <= 'z' {
		return true
	}
	return false
}
