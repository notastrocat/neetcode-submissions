const (
    RIGHT = iota
    LEFT
)

func largestRectangleArea(heights []int) int {
    maxArea := 0

    // these work as expected, kinda. they don't handle equals too well - but that's just my lack of understanding
    right := getSmallerElements(heights, RIGHT) // to get next smaller
    left := getSmallerElements(heights, LEFT) // to get prev smaller

	// fmt.Println(right)
	// fmt.Println(left)

    for i,v := range heights {
        rightBound := right[i]
        leftBound := left[i]

		// so we've got the exclusive boundaries... these are strictly to be excluded while
		// calculating the area.

        // If there is no smaller element to the right, the bound is exclusively the length of the array
		if rightBound == -1 {
			rightBound = len(heights)
		}
		// If leftBound == -1, we leave it! It acts perfectly as the exclusive boundary before index 0.

        area := v * (rightBound - leftBound - 1)

        if maxArea < area {
            maxArea = area
        }
    }

    return maxArea
}

func getSmallerElements(nums []int, direction int) []int {
	// this time i'm planning to store the index of the next smallest element on top of the stack instead of the real value.
	// let's see how this pans out.

    size := len(nums)
    res := make([]int, size)
    var stack []int

    switch direction {
    case RIGHT:
        for i := size - 1; i >= 0; i-- {
            for len(stack) != 0 && nums[i] <= nums[stack[len(stack)-1]] {
                // pop from stack
                stack = stack[:len(stack)-1]
            }

            if len(stack) == 0 {
                res[i] = -1
            } else {
                res[i] = stack[len(stack)-1]
            }

            stack = append(stack, i)
        }
    case LEFT:
        for i := 0; i < size; i++ {
            for len(stack) != 0 && nums[i] <= nums[stack[len(stack)-1]] {
                // pop from stack
                stack = stack[:len(stack)-1]
            }

            if len(stack) == 0 {
                res[i] = -1
            } else {
                res[i] = stack[len(stack)-1]
            }

            stack = append(stack, i)
        }
    default:
        fmt.Println("LOL!")
    }

    return res
}
