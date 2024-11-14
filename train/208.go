package train

type TrieNode struct {
	Count int
	Next  map[int32]*TrieNode
}

type Trie struct {
	Root *TrieNode
}

func Constructor208() Trie {
	return Trie{
		Root: &TrieNode{},
	}
}

func (this *Trie) Insert(word string) {
	cur := this.Root
	for _, c := range word {
		if cur.Next == nil {
			cur.Next = make(map[int32]*TrieNode)
			cur.Next[c] = &TrieNode{}
		} else {
			if _, exist := cur.Next[c]; !exist {
				cur.Next[c] = &TrieNode{}
			}
		}
		cur = cur.Next[c]
	}
	cur.Count++
}

func (this *Trie) Search(word string) bool {
	cur := this.Root
	for _, c := range word {
		if cur.Next == nil {
			return false
		} else {
			if _, exist := cur.Next[c]; !exist {
				return false
			}
		}
		cur = cur.Next[c]
	}
	return cur.Count > 0
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this.Root
	for _, c := range prefix {
		if cur.Next == nil {
			return false
		} else {
			if _, exist := cur.Next[c]; !exist {
				return false
			}
		}
		cur = cur.Next[c]
	}
	return true
}
