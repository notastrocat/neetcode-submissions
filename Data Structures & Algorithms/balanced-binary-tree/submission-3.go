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

func getHeight(root *TreeNode) int {
	if root == nil {
		return -1
	}

	rightHeight := getHeight(root.Right)
	leftHeight := getHeight(root.Left)

	return max(rightHeight, leftHeight) + 1
}

func isBalanced(root *TreeNode) bool {
	// a nil tree is always considered to be "balanced"
	if root == nil {
		return true
	}

    var stack []*TreeNode

	stack = append(stack, root)

	for len(stack) != 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// i'm a bum! always was. forgot to push children on to the "stack"
		if top.Right != nil {
			stack = append(stack, top.Right)
		}

		if top.Left != nil {
			stack = append(stack, top.Left)
		}

		rightHeight := getHeight(top.Right)
		leftHeight := getHeight(top.Left)

		diff := max(rightHeight - leftHeight, leftHeight - rightHeight)

		if diff > 1 {
			return false
		}
	}

	return true
}
