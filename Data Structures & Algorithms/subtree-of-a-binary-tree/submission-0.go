/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    } else if p == nil && q != nil {
        return false
    } else if p != nil && q == nil {
        return false
    }

    isRightSame := isSameTree(p.Right, q.Right)
    if !isRightSame {
        return false
    }

    isLeftSame := isSameTree(p.Left, q.Left)
    if !isLeftSame {
        return false
    }

    if q.Val != p.Val {
        return false
    }

    return isRightSame && isLeftSame
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    // probably this is redundant but i'll still do it.
    if root == nil && subRoot == nil {
        return true
    } else if root == nil && subRoot != nil {
        return false
    } else if root != nil && subRoot == nil {
        return false
    }

    var stack []*TreeNode
    stack = append(stack, root)

    for len(stack) != 0 {
        top := stack[len(stack)-1]
        stack = stack[:len(stack)-1]

        if top.Right != nil {
            stack = append(stack, top.Right)
        }

        if top.Left != nil {
            stack = append(stack, top.Left)
        }

        if isSameTree(top, subRoot) {
            return true
        }
    }

    return false
}
