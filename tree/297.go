package tree

import (
	"container/list"
	"strconv"
	"strings"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	queue := list.List{}
	queue.PushBack(root)
	var res string
	res += strconv.Itoa(root.Val)

	for queue.Len() != 0 {
		first := queue.Front()
		queue.Remove(first)
		node := first.Value.(*TreeNode)
		if node.Left != nil {
			res += "," + strconv.Itoa(node.Left.Val)
			queue.PushBack(node.Left)
		} else {
			res += ",null"
		}
		if node.Right != nil {
			res += "," + strconv.Itoa(node.Right.Val)
			queue.PushBack(node.Right)
		} else {
			res += ",null"
		}
	}
	return res
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	datas := strings.Split(data, ",")
	if len(datas) == 0 {
		return nil
	}
	if datas[0] == "null" || len(datas[0]) == 0 {
		return nil
	}
	val, _ := strconv.Atoi(datas[0])
	queue := list.List{}

	root := &TreeNode{Val: val}
	queue.PushBack(root)

	curIndex := 0
	for queue.Len() > 0 {
		front := queue.Front()
		queue.Remove(front)

		curIndex++
		treeNode := front.Value.(*TreeNode)
		if datas[curIndex] == "null" {
			treeNode.Left = nil
		} else {
			val, _ = strconv.Atoi(datas[curIndex])
			leftNode := &TreeNode{Val: val}
			treeNode.Left = leftNode
			queue.PushBack(leftNode)
		}

		curIndex++
		if datas[curIndex] == "null" {
			treeNode.Right = nil
		} else {
			val, _ = strconv.Atoi(datas[curIndex])
			rightNode := &TreeNode{Val: val}
			treeNode.Right = rightNode
			queue.PushBack(rightNode)
		}
	}
	return root
}
