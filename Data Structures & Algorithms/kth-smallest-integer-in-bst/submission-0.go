/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// yes, call  me dumb. but i'm gonna submit this.
func helper_flatten_tree(root *TreeNode, list []int) []int {
	if root == nil {
		return list
	}

	list = helper_flatten_tree(root.Left, list)
	list = append(list, root.Val)
	list = helper_flatten_tree(root.Right, list)

	return list
}

func kthSmallest(root *TreeNode, k int) int  {
	var tree_nodes []int
	tree_nodes = helper_flatten_tree(root, tree_nodes)

	return tree_nodes[k-1]
}
