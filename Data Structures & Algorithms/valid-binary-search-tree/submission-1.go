/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// i'm so fucking dumb that i'll do this and call it a day.
// lol, not really - will check if there are better ways to do this. there must be.
func helper_flatten_tree(root *TreeNode, list []int) []int {
	if root == nil {
		return list
	}

	list = helper_flatten_tree(root.Left, list)
	list = append(list, root.Val)
	list = helper_flatten_tree(root.Right, list)

	return list
}

func isValidBST(root *TreeNode) bool {
	var tree_nodes []int
	tree_nodes = helper_flatten_tree(root, tree_nodes)

	// fmt.Println(tree_nodes)

	for i:=0; i<len(tree_nodes)-1; i++ {
		if tree_nodes[i] >= tree_nodes[i+1] {
			return false
		}
	}

	return true
}