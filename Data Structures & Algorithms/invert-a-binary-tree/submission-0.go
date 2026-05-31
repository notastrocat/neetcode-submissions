/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
		return nil
	}

	rightTree := invertTree(root.Right)
	leftTree := invertTree(root.Left)

	root.Right, root.Left = leftTree, rightTree

	return root
}
