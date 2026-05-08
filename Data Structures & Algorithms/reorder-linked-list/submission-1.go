/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
    if head == nil || head.Next == nil || head.Next.Next == nil {
        return
    }

    middle := findMiddle(head)

    secondHalf := reverseList(middle.Next) 

    middle.Next = nil 

    mergeLists(head, secondHalf)
}

// findMiddle returns the middle node of a linked list using the tortoise and hare approach.
func findMiddle(head *ListNode) *ListNode {
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    return slow
}

func reverseList(head *ListNode) *ListNode {
    var prev *ListNode = nil
    curr := head

    for curr != nil {
        nextNode := curr.Next
        curr.Next = prev
        prev = curr
        curr = nextNode
    }

    return prev
}

func mergeLists(l1, l2 *ListNode) {
    for l2 != nil {
        tmp1 := l1.Next
        tmp2 := l2.Next

        l1.Next = l2
        l2.Next = tmp1

        l1 = tmp1
        l2 = tmp2
    }
}