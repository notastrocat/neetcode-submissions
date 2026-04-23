func isBalanced(root *TreeNode) bool {
    return checkHeight(root) != -1
}

func checkHeight(node *TreeNode) int {
    if node == nil {
        return 0
    }

    leftHeight := checkHeight(node.Left)
    if leftHeight == -1 {
        return -1
    }

    rightHeight := checkHeight(node.Right)
    if rightHeight == -1 {
        return -1
    }

    if absDiff(leftHeight, rightHeight) > 1 {
        return -1
    }

    // return the actual height of this node; not diff+1
    return max(leftHeight, rightHeight) + 1
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func absDiff(a, b int) int {
    if a > b {
        return a - b
    }
    return b - a
}