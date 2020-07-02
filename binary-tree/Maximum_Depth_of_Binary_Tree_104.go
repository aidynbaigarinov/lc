package btree

func maxDepth(root *TreeNode) int {
	return helper(root, 0)
}

func helper(root *TreeNode, depth int) int {
	if root == nil {
		return depth
	}
	depth++
	left := helper(root.Left, depth)
	right := helper(root.Right, depth)

	if left > depth && left >= right {
		depth = left
	} else if right > depth && right > left {
		depth = right
	}
	return depth
}