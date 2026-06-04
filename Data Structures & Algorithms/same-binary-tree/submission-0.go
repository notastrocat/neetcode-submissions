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
