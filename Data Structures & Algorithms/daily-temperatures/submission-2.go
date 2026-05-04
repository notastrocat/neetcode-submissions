func dailyTemperatures(temperatures []int) []int {
	maxLen := len(temperatures)
	res := make([]int, maxLen)
	var stack []int

	for i,v := range temperatures {
		if len(stack) == 0 {
			stack = append(stack, i)
			continue
		}

		for len(stack) != 0 && v > temperatures[stack[len(stack)-1]] {
			// pop from stack
			last := len(stack)-1
			top := stack[last]
			stack = stack[:last]

			// store the diff
			res[top] = i - top
		}

		stack = append (stack, i)
	}

	for len(stack) != 0 {
		last := len(stack)-1
		top := stack[last]
		stack = stack[:last]

		res[top] = 0
	}

	return res
}

// very naïve way to solve would be to iterate over and over.
// [30,38,30,36,35,40,28]
// [ 0, 1, 2, 3, 4, 5, 6 ]

// [ 1, 4, 1, 2, 1, 0, 0 ]