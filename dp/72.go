package dp

func minDistance(word1 string, word2 string) int {
	word1Len := len(word1)
	word2Len := len(word2)
	var f = make([][]int, word1Len+1)
	for i := range f {
		f[i] = make([]int, word2Len+1)
		for j := range f[i] {
			f[i][0] = i
			f[0][j] = j
		}
	}

	for i := 1; i <= word1Len; i++ {
		for j := 1; j <= word2Len; j++ {
			if word1[i-1] == word2[j-1] {
				f[i][j] = f[i-1][j-1]
			} else {
				f[i][j] = min(f[i][j-1], min(f[i-1][j], f[i-1][j-1])) + 1
			}
		}
	}
	return f[word1Len][word2Len]
}
