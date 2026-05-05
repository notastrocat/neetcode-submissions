func nextGreaterElement(nums1 []int, nums2 []int) []int {
	size := len(nums2)
	var stack []int
	res := make(map[int]int, size)

	finalRes := make([]int, len(nums1))

	stack = append(stack, nums2[size-1])
	res[nums2[size-1]] = -1

	for i:=size-2; i>=0; i-- {
		for len(stack) != 0 && nums2[i] >= stack[len(stack)-1] {
			//pop from stack
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			res[nums2[i]] = -1
		} else {
			res[nums2[i]] = stack[len(stack)-1]
		}

		stack = append(stack, nums2[i])
	}

	for i, v := range nums1 {
		finalRes[i] = res[v]
	}

	return finalRes
}
