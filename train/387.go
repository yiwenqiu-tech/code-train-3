package train

import "container/list"

func firstUniqChar(s string) int {
	l := list.New()
	m := make(map[int32]*list.Element)
	for index, c := range s {
		if item, exist := m[c]; exist {
			m[c] = nil
			if item != nil {
				l.Remove(item)
			}
		} else {
			l.PushBack(index)
			m[c] = l.Back()
		}
	}
	if l.Len() == 0 {
		return -1
	} else {
		return l.Front().Value.(int)
	}
}
