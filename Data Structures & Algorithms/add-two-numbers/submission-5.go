/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// yeah so the initial solution was bullshit where I was
// reading the list into two numbers and then adding
// and then creating another list from the result.

// technically, it would only work in cases the numbers
// are within 64-bit range. but that is not the case here.

// so the last submission attempts to build up on the idea
// to make use of the first list as the result list.
// now, only main problem to handle is that when the two
// lists are of different lengths. otherwise, getting the
// carry-over digit and storing it and making it zero
// again when done with it is alright.

// so far i am treating the second list as something
// where we keep on reading until we reach the end.
// then, in case there are still carry-overs we keep
// on adding it to the last digit(s) of the first list.

// and of course, all along we keep track if the first
// list has reached its end, if so => we create an extra
// last node. and then later on, delete it before returning
// the head of the result list, i.e., the first list.

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if nil == l1 && nil != l2 {
		return l2
	} else if nil != l1 && nil == l2 {
		return l1
	} else if nil == l1 && nil == l2 {
		return nil
	}

    ptr1 := l1
	ptr2 := l2
	lastPtr := l1
	carry := 0

	for ptr2 != nil {
		ptr1.Val += ptr2.Val + carry

		if 0 != ( ptr1.Val / 10 ) {
			// this means that carry should be set to the
			// extra digit in the ten's place
			digit := ptr1.Val % 10
			carry = ptr1.Val / 10

			ptr1.Val = digit
		} else {
			carry = 0
		}

		lastPtr = ptr1

		if nil == ptr1.Next {
			// create a node since list 1 is exhausted
			newNode := &ListNode{
				Val: 0,
				Next: nil,
			}
			ptr1.Next = newNode
		}

		ptr1 = ptr1.Next
		ptr2 = ptr2.Next
	}

	for 0 != carry {
		ptr1.Val += carry

		if 0 != ( ptr1.Val / 10 ) {
			// this means that carry should be set to the
			// extra digit in the ten's place
			digit := ptr1.Val % 10
			carry = ptr1.Val / 10

			ptr1.Val = digit
		} else {
			carry = 0
		}

		lastPtr = ptr1
		if nil == ptr1.Next {
			// create a node since list 1 is exhausted
			newNode := &ListNode{
				Val: 0,
				Next: nil,
			}
			ptr1.Next = newNode
		}

		ptr1 = ptr1.Next
	}

	if lastPtr.Next.Val == 0 {
		// fmt.Println("this case needs to be handled.")
		lastPtr.Next = nil
	}

	return l1
}
