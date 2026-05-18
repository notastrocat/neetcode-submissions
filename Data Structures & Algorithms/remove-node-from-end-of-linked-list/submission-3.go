/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{Next: head}

    length := 0
    current := head
    for current != nil {
        length++
        current = current.Next
    }

    targetIndex := length - n

    // Start at dummy so if targetIndex is 0, we stay on dummy (perfect for deleting head)
    current = dummy
    for i := 0; i < targetIndex; i++ {
        current = current.Next
    }

    current.Next = current.Next.Next

    return dummy.Next
}