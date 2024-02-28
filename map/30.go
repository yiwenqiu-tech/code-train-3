package _map

func findSubstring(s string, words []string) []int {
	if len(words) < 1 {
		return nil
	}
	wordsSize := len(words)
	wordLen := len(words[0])
	totalWordLen := wordLen * wordsSize
	if len(s) < totalWordLen { // s的长度还没有totalWordLen长，直接返回
		return nil
	}
	var wordsMap = make(map[string]int)
	for _, w := range words {
		wordsMap[w]++
	}

	var result []int
	for i := 0; i < len(s); i++ { // TODO 存在优化空间，可以分组处理
		if i+totalWordLen > len(s) { // 后面的不够长了，就不比较了
			break
		}
		tmp := s[i : i+totalWordLen]
		if isEqual(tmp, wordLen, wordsMap) {
			result = append(result, i)
		}
	}
	return result
}

func isEqual(s string, wordsLen int, wordsMap map[string]int) bool {
	swordsMap := make(map[string]int)
	for i := 0; i < len(s); i += wordsLen {
		swordsMap[s[i:i+wordsLen]]++
	}
	for w, count := range wordsMap {
		c, exist := swordsMap[w]
		if !exist {
			return false
		}
		if c != count {
			return false
		}
	}
	return true
}
