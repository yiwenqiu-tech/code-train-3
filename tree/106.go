package tree

var postOrder2 []int
var inOrder2 []int
var inOrderMap2 map[int]int

func buildTree2(inorder []int, postorder []int) *TreeNode {
	inOrder2 = inorder
	postOrder2 = postorder
	inOrderMap2 = make(map[int]int)
	for index, i := range inOrder2 {
		inOrderMap2[i] = index
	}
	return recurForBuildTree2(0, len(inorder)-1, 0, len(postorder)-1)
}

func recurForBuildTree2(il1, ir1, pl1, pr1 int) *TreeNode {
	if il1 > ir1 || pl1 > pr1 {
		return nil
	}
	rootVal := postOrder2[pr1] // 后序遍历最后一个数为根节点
	inOrderIndex := inOrderMap2[rootVal]

	node := &TreeNode{Val: rootVal}

	node.Left = recurForBuildTree2(il1, inOrderIndex-1, pl1, pl1+inOrderIndex-1-il1)

	node.Right = recurForBuildTree2(inOrderIndex+1, ir1, pr1-1-(ir1-(inOrderIndex+1)), pr1-1)

	return node
}
