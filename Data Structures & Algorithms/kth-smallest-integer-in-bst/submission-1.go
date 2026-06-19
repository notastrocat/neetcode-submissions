/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func helper_fn(root *TreeNode, cnt *int, res *int) {
	if root == nil {
		return
	}

	helper_fn(root.Left, cnt, res)

    if *cnt == 0 {
        return
    }

    *(cnt)--

    if *cnt == 0 {
        *res = root.Val 
        return
    }

	helper_fn(root.Right, cnt, res)

	// return cnt
}

func kthSmallest(root *TreeNode, k int) int  {
	var cnt int = k
    var res int
	helper_fn(root, &cnt, &res)

	return res
}
