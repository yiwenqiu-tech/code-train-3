package tree

import (
	"container/list"
)

type levelOrderVal struct {
	level int
	val   *Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	queue := list.List{}
	queue.PushBack(&levelOrderVal{
		level: 1,
		val:   root,
	})
	var res [][]int
	for queue.Len() != 0 {
		head := queue.Front()
		queue.Remove(head)
		l := head.Value.(*levelOrderVal).level
		v := head.Value.(*levelOrderVal).val
		if len(res) < l { // 若当前层不存在，则添加一个元素
			res = append(res, []int{})
		}
		res[l-1] = append(res[l-1], v.Val)
		for _, c := range v.Children {
			queue.PushBack(&levelOrderVal{
				level: l + 1,
				val:   c,
			})
		}
	}
	return res
}
