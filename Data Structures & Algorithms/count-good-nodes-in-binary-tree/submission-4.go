/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// ez pz. but having a global counter fucks things up since this if Golang and the test suite is
// part of the "main" package. so a package-wide global counter isn't thread-safe.
func helper_dfs(root *TreeNode, maxSoFar int, goodNodesCount *int) {
	if root == nil {
		return
	}

	if maxSoFar <= root.Val {
		maxSoFar = root.Val
		*(goodNodesCount)++
	}

	helper_dfs(root.Right, maxSoFar, goodNodesCount)
	helper_dfs(root.Left, maxSoFar, goodNodesCount)
}

func goodNodes(root *TreeNode) int {
    var goodNodesCount int = 0

	helper_dfs(root, math.MinInt16, &goodNodesCount)

	return goodNodesCount
}
