package train

type WordDictionary struct {
	trie Trie
}

func Constructor() WordDictionary {
	return WordDictionary{
		trie: Constructor208(),
	}
}

func (this *WordDictionary) AddWord(word string) {
	this.trie.Insert(word)
}

func (this *WordDictionary) Search(word string) bool {
	return this.dfs([]byte(word), 0, this.trie.Root)
}

func (this *WordDictionary) dfs(bs []byte, cur int, trieNode *TrieNode) bool {
	if cur == len(bs) {
		return trieNode.Count > 0
	}
	if trieNode.Next == nil {
		return false
	}
	if isLowerLetter(bs[cur]) {
		if elem, exist := trieNode.Next[int32(bs[cur])]; !exist {
			return false
		} else {
			return this.dfs(bs, cur+1, elem)
		}
	} else {
		for _, elem := range trieNode.Next {
			if this.dfs(bs, cur+1, elem) {
				return true
			}
		}
	}
	return false
}

func isLowerLetter(c byte) bool {
	if c >= 'a' && c <= 'z' {
		return true
	}
	return false
}
