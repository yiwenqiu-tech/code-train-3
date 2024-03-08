package search

import "container/list"

var validChar = []byte{'A', 'C', 'G', 'T'}

func minMutation(startGene string, endGene string, bank []string) int {
	queue := list.List{}
	var geneLevel = make(map[string]int)

	queue.PushBack(startGene)
	geneLevel[startGene] = 0

	bankMap := make(map[string]struct{})
	for _, b := range bank {
		bankMap[b] = struct{}{}
	}

	for queue.Len() != 0 {
		head := queue.Front()
		queue.Remove(head)
		headVal := head.Value.(string)
		curLevel := geneLevel[headVal]
		for index := range headVal {
			for _, c := range validChar {
				if c == headVal[index] { // 当前字符已经是了，直接提过不处理
					continue
				}
				strBytes := []byte(headVal)
				strBytes[index] = c // 替换字符
				changeStr := string(strBytes)
				if _, exist := bankMap[changeStr]; !exist { // 判断替换的字符是否在bank里，不存在则为非法基因，直接跳过
					continue
				}
				if _, exist := geneLevel[changeStr]; exist { // 已经遍历过，则跳过
					continue
				}
				geneLevel[changeStr] = curLevel + 1
				queue.PushBack(changeStr)
				if changeStr == endGene {
					return curLevel + 1
				}
			}
		}
	}
	return -1
}
