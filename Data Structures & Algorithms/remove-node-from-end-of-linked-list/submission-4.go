/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{
        Next: head,
    }

    slow := dummy
    fast := dummy

    // <= is imp here since we need a gap of n not n-1.
    for i:=0; i<=n; i++ {
        fast = fast.Next
    }

    // till fast is valid we move both pointers.
    for fast != nil {
        slow = slow.Next
        fast = fast.Next
    }

    slow.Next = slow.Next.Next

    return dummy.Next
}