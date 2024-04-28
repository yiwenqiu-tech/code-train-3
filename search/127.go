package search

import "container/list"

func ladderLength(beginWord string, endWord string, wordList []string) int {
	var wordMap = make(map[string]struct{})
	for i := range wordList {
		wordMap[wordList[i]] = struct{}{}
	}
	if _, exist := wordMap[endWord]; !exist { // 终点都不在wordList，则直接返回即可。
		return 0
	}
	wordLen := len(beginWord)
	var depthMap = make(map[string]int)
	queue := list.New()
	queue.PushBack(beginWord)
	depthMap[beginWord] = 1

	for queue.Len() != 0 {
		frontItem := queue.Front()
		queue.Remove(frontItem)
		frontItemString := frontItem.Value.(string)
		frontItemLevel := depthMap[frontItemString]
		for i := 0; i < wordLen; i++ {
			tmp := []byte(frontItemString)
			for j := 0; j < 26; j++ {
				if frontItemString[i] == byte('a'+j) { // 当前已经为j对应的字符了
					continue
				}
				tmp[i] = byte('a' + j)
				if string(tmp) == endWord {
					return frontItemLevel + 1
				}
				if _, exist := wordMap[string(tmp)]; !exist { // 是否在wordList，不在的话不要
					continue
				}
				if _, exist := depthMap[string(tmp)]; exist { // 是否已经遍历过了，是的话也不动了
					continue
				}
				depthMap[string(tmp)] = frontItemLevel + 1
				queue.PushBack(string(tmp))
			}
		}
	}
	return 0
}

func ladderLength2(beginWord string, endWord string, wordList []string) int {
	var wordMap = make(map[string]struct{})
	for i := range wordList {
		wordMap[wordList[i]] = struct{}{}
	}
	if _, exist := wordMap[endWord]; !exist { // 终点都不在wordList，则直接返回即可。
		return 0
	}
	wordLen := len(beginWord)

	var beginDepthMap = make(map[string]int)
	beginQueue := list.New()
	beginQueue.PushBack(beginWord)
	beginDepthMap[beginWord] = 1

	var endDepthMap = make(map[string]int)
	endQueue := list.New()
	endQueue.PushBack(endWord)
	endDepthMap[endWord] = 1

	for beginQueue.Len() != 0 || endQueue.Len() != 0 {
		beginQueueSize := beginQueue.Len() // 每一次扩展，得把当前层遍历完，只扩展一个会把两层数据混在一起，导致数据不对
		for beginQueueSize > 0 {
			frontItem := beginQueue.Front()
			beginQueue.Remove(frontItem)
			beginQueueSize--
			frontItemString := frontItem.Value.(string)
			frontItemLevel := beginDepthMap[frontItemString]
			for i := 0; i < wordLen; i++ {
				tmp := []byte(frontItemString)
				for j := 0; j < 26; j++ {
					if frontItemString[i] == byte('a'+j) { // 当前已经为j对应的字符了
						continue
					}
					tmp[i] = byte('a' + j)
					if value, exist := endDepthMap[string(tmp)]; exist { // 相遇了
						return frontItemLevel + value
					}
					if _, exist := wordMap[string(tmp)]; !exist { // 是否在wordList，不在的话不要
						continue
					}
					if _, exist := beginDepthMap[string(tmp)]; exist { // 是否已经遍历过了，是的话也不动了
						continue
					}
					beginDepthMap[string(tmp)] = frontItemLevel + 1
					beginQueue.PushBack(string(tmp))
				}
			}
		}
		endQueueSize := endQueue.Len()
		for endQueueSize > 0 { // 每一次扩展，得把当前层遍历完，只扩展一个会把两层数据混在一起，导致数据不对
			frontItem := endQueue.Front()
			endQueue.Remove(frontItem)
			endQueueSize--
			frontItemString := frontItem.Value.(string)
			frontItemLevel := endDepthMap[frontItemString]
			for i := 0; i < wordLen; i++ {
				tmp := []byte(frontItemString)
				for j := 0; j < 26; j++ {
					if frontItemString[i] == byte('a'+j) { // 当前已经为j对应的字符了
						continue
					}
					tmp[i] = byte('a' + j)
					if value, exist := beginDepthMap[string(tmp)]; exist { // 相遇了
						return frontItemLevel + value
					}
					if _, exist := wordMap[string(tmp)]; !exist { // 是否在wordList，不在的话不要
						continue
					}
					if _, exist := endDepthMap[string(tmp)]; exist { // 是否已经遍历过了，是的话也不动了
						continue
					}
					endDepthMap[string(tmp)] = frontItemLevel + 1
					endQueue.PushBack(string(tmp))
				}
			}
		}
	}
	return 0
}
