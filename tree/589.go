package tree

type Node struct {
	Val      int
	Children []*Node
}

var resForPreOrder []int

func preorder(root *Node) []int {
	resForPreOrder = nil
	recurForPreOrder(root)
	return resForPreOrder
}

func recurForPreOrder(node *Node) {
	if node == nil {
		return
	}
	resForPreOrder = append(resForPreOrder, node.Val)
	for _, n := range node.Children {
		recurForPreOrder(n)
	}
}
