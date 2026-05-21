// my r*t*rded ass took hours to realize the fact that since the numbers in the array
// are all in the range [1,n]... each number is a valid index of the array (because array is n+1).
// so i can treat the array as a linked list where number in the array is supposed to be the
// pointer to the *next*. next = nums[index]

// how this helps? well, think about it... if there are duplicates, the nums[index] (which is next
// node location) will point to the same node for two (or more) times. voila! we have a loop in
// the linked list. and we know how to detect loops in linked lists.

func findDuplicate(nums []int) int {
    slow := nums[0]
	fast := nums[0]

	for {
		slow = nums[slow] // move slow once
		fast = nums[nums[fast]] // move fast twice

		// you see, in this case... it might be difficult to wrap head around the fact that
		// where are we moving to now? think - write a case on paper.

		if slow == fast {
			break
		}
	}

	//now that we've detected that there *is* a loop in the linked list... what to do?
	// how to proceed now??? 🤔

	slow = nums[0] // reset and move once each pointer... i think that they'll collide since fast is now
	// inside a cycle. i mean that's how we used to find the point of cycle.

	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}
