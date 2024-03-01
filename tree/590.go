package tree

var resForPostOrder []int

func postorder(root *Node) []int {
	resForPostOrder = nil
	recurForPostOrder(root)
	return resForPostOrder
}

func recurForPostOrder(node *Node) {
	if node == nil {
		return
	}
	for _, n := range node.Children {
		recurForPostOrder(n)
	}
	resForPostOrder = append(resForPostOrder, node.Val)
}
