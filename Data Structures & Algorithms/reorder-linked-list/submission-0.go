/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
    // let me try to do something ez

	size := 0
	mover := head
	for mover != nil {
		mover = mover.Next
		size++
	}

	ptr := head

	var i int
	if size%2 == 0 {
		i = (size/2)-1
	} else {
		i = int(size/2)
	}

	for i!=0 {
		it := ptr.Next
		for it.Next.Next != nil {
			it = it.Next
		}

		tmp := it.Next

		// cut ties
		it.Next = nil

		// insert
		tmp.Next = ptr.Next

		// make connection
		ptr.Next = tmp

		ptr = ptr.Next.Next
		i--
	}
}
