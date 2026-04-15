/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
    if  head == nil {
        return nil
    }

    if head.Next == nil {
        return head
    }

    newHead := reverseList(head.Next)
    // cur := head

    // head = newHead
    // for (head.Next != nil) {
    //     head = head.Next
    // }
    // head.Next = cur
    // cur.Next = nil

    head.Next.Next = head
    head.Next = nil

    return newHead
}
