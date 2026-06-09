/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func doSomething(root, node *TreeNode, list []*TreeNode) []*TreeNode {
	if root == nil {
		fmt.Println("fucking passing a nil node!")
		return list
	}

	var ptr = root
	for ptr != nil {
		list = append(list, ptr)
		if ptr.Val == node.Val {
			return list
		} else if ptr.Val < node.Val {
			ptr = ptr.Right
		} else {
			ptr = ptr.Left
		}
	}

	return list
}

// i think i am missing some nil checks and/or some array out of bounds check...

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	var pList []*TreeNode
	var qList []*TreeNode

	pList = doSomething(root, p, pList)
	qList = doSomething(root, q, qList)

	shorter := min(len(pList), len(qList))

	retVal := root
	for i := range shorter {
		if pList[i].Val == qList[i].Val {
			retVal = pList[i]
		} else {
			break
		}
	}

	return retVal
}
