/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var res [][]int
	var tmp []int

	var queue []*TreeNode
	queue = append(queue, root)
	queue = append(queue, nil)

	for len(queue) != 0 {
		front := queue[0]
		queue = queue[1:]

		if front != nil {
			// fmt.Printf("%d", front.Val)
			tmp = append(tmp, front.Val)
		}

		if front == nil {
			if len(queue) == 0 {
				res = append(res, tmp)
				break
			}

			queue = append(queue, nil)
			// fmt.Println()
			res = append(res, tmp)
			// is this the way to clear the contents of an array in GOlang?
			tmp = []int{}
			continue
		}

		// idk if order matters, i'll still have it in case it matters.
		if front.Left != nil {
			queue = append(queue, front.Left)
		}
		if front.Right != nil {
			queue = append(queue, front.Right)
		}
	}

	return res
}

// i don't even know the complexity for this one...

func rightSideView(root *TreeNode) []int {
    levelOrder := levelOrder(root)

	var res []int

	for _,level := range levelOrder {
		if len(level) > 0 {
			res = append(res, level[len(level)-1])
		}
	}

	return res
}
















