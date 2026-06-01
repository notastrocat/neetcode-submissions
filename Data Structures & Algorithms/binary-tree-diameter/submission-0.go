/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	rightDepth := maxDepth(root.Right)
	leftDepth := maxDepth(root.Left)

	return max(rightDepth, leftDepth) + 1
}

// bs! this was supposed to be ez. i am just plain dumb. came *this* close to the solution
// well, the naive solution at least. was fucking up the summing of the left and right heights.
// didn't read the question properly, then mis-interpreted the hint.
// anyway, imma optimize later. got butt-load of stuff to do at work rn.
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var stack []*TreeNode
	stack = append(stack, root)
	dia := 0

	for len(stack) !=  0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if top.Right != nil {
			stack = append(stack, top.Right)
		}
		if top.Left != nil {
			stack = append(stack, top.Left)
		}

		leftDepth := maxDepth(top.Left)
        rightDepth := maxDepth(top.Right)
        localDiameter := leftDepth + rightDepth
		if localDiameter > dia {
			dia = localDiameter
		}
	}

	return dia
}
