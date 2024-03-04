package tree

import (
	"strconv"
	"strings"
)

type CodecA struct {
	serializeStrings []string
	deserializeCur   int
}

func ConstructorA() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *CodecA) serialize(root *TreeNode) string {
	this.recurForSerialize(root)
	return strings.Join(this.serializeStrings, ",")
}

func (this *CodecA) recurForSerialize(node *TreeNode) {
	if node == nil {
		this.serializeStrings = append(this.serializeStrings, "N")
		return
	}
	this.serializeStrings = append(this.serializeStrings, strconv.Itoa(node.Val))
	this.recurForSerialize(node.Left)
	this.recurForSerialize(node.Right)
}

// Deserializes your encoded data to tree.
func (this *CodecA) deserialize(data string) *TreeNode {
	this.deserializeCur = 0
	datas := strings.Split(data, ",")
	return this.recurForDeserialize(datas)
}

func (this *CodecA) recurForDeserialize(datas []string) *TreeNode {
	if len(datas) == 0 {
		return nil
	}
	data := datas[this.deserializeCur]
	if data == "N" {
		this.deserializeCur++
		return nil
	}
	val, _ := strconv.Atoi(data)
	treeNode := &TreeNode{
		Val: val,
	}
	this.deserializeCur++
	treeNode.Left = this.recurForDeserialize(datas)
	treeNode.Right = this.recurForDeserialize(datas)
	return treeNode
}
