/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// i'll attempt to build the height logic within the same fn call to reduce overall
// complexity of the code... later.

func helper(root *TreeNode) (int, bool) {
    if root == nil {
        return 0, true
    }

    leftHeight, leftBalanced := helper(root.Left)
    if !leftBalanced {
        return 0, false
    }

    rightHeight, rightBalanced := helper(root.Right)
    if !rightBalanced {
        return 0, false
    }

    diff := max(leftHeight-rightHeight, rightHeight-leftHeight)
    if diff > 1 {
        return 0, false
    }

    return max(leftHeight, rightHeight) + 1, true
}

func isBalanced(root *TreeNode) bool {
    _, balanced := helper(root)
    return balanced
}
