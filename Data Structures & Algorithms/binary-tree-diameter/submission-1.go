/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := math.MinInt32
	calculateHeight(root, &maxDiameter)
	return maxDiameter
}

func calculateHeight(node *TreeNode, maxDiameter *int) int {
	if node == nil {
		return -1
	}

	leftHeight := calculateHeight(node.Left, maxDiameter)
	rightHeight := calculateHeight(node.Right, maxDiameter)

	// adjust the two -1s
	currentDiameter := leftHeight + rightHeight + 2

	if currentDiameter > *maxDiameter {
		*maxDiameter = currentDiameter
	}

	// we're actually calculating the height so do the +1 trick here as well, I removed it
	// classic dumb me. so naturally it was just a big ole ).
	return max(leftHeight, rightHeight) + 1
}
