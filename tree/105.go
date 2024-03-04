package tree

var preOrder []int
var inOrder []int
var inOrderMap map[int]int

func buildTree(preorder []int, inorder []int) *TreeNode {
	preOrder = preorder
	inOrder = inorder

	inOrderMap = make(map[int]int)
	for index, i := range inOrder {
		inOrderMap[i] = index
	}

	return recurForBuildTree(0, len(preorder)-1, 0, len(inorder)-1)
}

func recurForBuildTree(pl1, pr1, il1, ir1 int) *TreeNode {
	if pl1 > pr1 || il1 > ir1 {
		return nil
	}
	root := &TreeNode{Val: preOrder[pl1]}

	inOrderIndex := inOrderMap[preOrder[pl1]] // 在inorder的序号

	root.Left = recurForBuildTree(pl1+1, pl1+1+(inOrderIndex-1-il1), il1, inOrderIndex-1)

	root.Right = recurForBuildTree(pl1+1+(inOrderIndex-1-il1)+1, pr1, inOrderIndex+1, ir1)

	return root
}
