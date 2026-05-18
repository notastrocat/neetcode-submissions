/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if nil == head {
		return head
	}

    ptr := head
	var nums []int
	for ptr != nil {
		nums = append(nums, ptr.Val)
		ptr = ptr.Next
	}

	size := len(nums)

	if 1 == size {
		head = nil
		return head
	}

	// assume 0-based index now
	idx := size - n
	cur := head

	if 0 == idx {
		head = cur.Next
		return head
	}

	for i:=0; i<idx-1; i++ {
		cur = cur.Next
	}

	del := cur.Next
	cur.Next = cur.Next.Next
	del.Next = nil

	return head
}
