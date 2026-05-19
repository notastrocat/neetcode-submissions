/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	newListHead := &Node{
		Val: head.Val,
		Next: nil,
		Random: nil,
	}

	table := make(map[*Node]*Node)

	table[head] = newListHead

	// just build the new linked list for now.
	ptr := head.Next
	newPtr := newListHead

	for ptr != nil {
		newNode := &Node{
			Val: ptr.Val,
			Next: nil,
			Random: nil,
		}
		table[ptr] = newNode

		newPtr.Next = newNode

		ptr = ptr.Next
		newPtr = newPtr.Next
	}

	// fmt.Println(table)

	ptr = head
	newPtr = newListHead
	// now, connect the random pointers of the new list, based on the old list.
	for ptr != nil {
		if copyLocation, ok := table[ptr.Random]; ok {
			newPtr.Random = copyLocation
		} else {
			newPtr.Random = nil
		}

		ptr = ptr.Next
		newPtr = newPtr.Next
	}

	return newListHead
}
