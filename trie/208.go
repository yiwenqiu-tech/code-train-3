package trie

type TrieNode struct {
	Count int
	Next  map[byte]*TrieNode
}

type Trie struct {
	Root *TrieNode
}

func Constructor() Trie {
	return Trie{
		Root: &TrieNode{},
	}
}

func (this *Trie) Insert(word string) {
	//var cur = this.Root
	//for i := range word {
	//	index := word[i] - 'a'
	//	if cur.Next[index] == nil {
	//		cur.Next[index] = &TrieNode{}
	//	}
	//	cur = cur.Next[index]
	//}
	//cur.Count++

	var cur = this.Root
	for i := range word {
		if cur.Next == nil {
			cur.Next = make(map[byte]*TrieNode)
			cur.Next[word[i]] = &TrieNode{}
		} else if _, exist := cur.Next[word[i]]; !exist {
			cur.Next[word[i]] = &TrieNode{}
		}
		cur = cur.Next[word[i]]
	}
	cur.Count++
}

func (this *Trie) Search(word string) bool {
	//var cur = this.Root
	//for i := range word {
	//	index := word[i] - 'a'
	//	if cur.Next[index] == nil {
	//		return false
	//	}
	//	cur = cur.Next[index]
	//}
	//if cur.Count > 0 {
	//	return true
	//}
	//return false

	var cur = this.Root
	for i := range word {
		if cur.Next == nil {
			return false
		} else if _, exist := cur.Next[word[i]]; !exist {
			return false
		}
		cur = cur.Next[word[i]]
	}
	if cur.Count > 0 {
		return true
	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	//var cur = this.Root
	//for i := range prefix {
	//	index := prefix[i] - 'a'
	//	if cur.Next[index] == nil {
	//		return false
	//	}
	//	cur = cur.Next[index]
	//}
	//return true

	var cur = this.Root
	for i := range prefix {
		if cur.Next == nil {
			return false
		} else if _, exist := cur.Next[prefix[i]]; !exist {
			return false
		}
		cur = cur.Next[prefix[i]]
	}
	return true
}
