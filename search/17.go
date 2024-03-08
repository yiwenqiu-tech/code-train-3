package search

var input string

var letterCombinationsRes []string

var digitMap = map[byte][]byte{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	letterCombinationsRes = nil
	if len(digits) == 0 {
		return nil
	}
	input = digits
	dfsForLetterCombinations(nil)
	return letterCombinationsRes
}

func dfsForLetterCombinations(curStr []byte) {
	if len(curStr) == len(input) {
		letterCombinationsRes = append(letterCombinationsRes, string(curStr))
		return
	}
	index := len(curStr)
	num := input[index]
	digitOfNum := digitMap[num]
	for _, d := range digitOfNum {
		curStr = append(curStr, d)
		dfsForLetterCombinations(curStr)
		curStr = curStr[:len(curStr)-1]
	}
}
