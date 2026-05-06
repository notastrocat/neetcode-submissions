const (
    RIGHT = iota
    LEFT
)

func largestRectangleArea(heights []int) int {
    maxArea := 0

    // these work as expected, kinda. they don't handle equals too well - but that's just my lack of understanding
    right := getSmallerElements(heights, RIGHT) // to get next smaller
    left := getSmallerElements(heights, LEFT) // to get prev smaller

    for i,v := range heights {
        rightBound := right[i]
        leftBound := left[i]

        // var width int
        // var height int

        if -1 == rightBound {
            // there is no next element smaller than cur i
            rightBound = len(heights)-1
        }
        if -1 == leftBound {
            // there is no prev element smaller than cur i
            leftBound = 0
        }

        area := v * (rightBound - leftBound + 1)

        if maxArea < area {
            maxArea = area
        }
    }

    return maxArea
}

func getSmallerElements(nums []int, direction int) []int {
	// this time i'm planning to store the index of the next smallest element on top of the stack instead of the real value.
	// let's see how this pans out.

    // i also need to look for a way to handle equal values.

    size := len(nums)
    res := make([]int, size)
    var stack []int

    switch direction {
    case RIGHT:
        res[size-1] = -1
        stack = append(stack, size-1)

        for i:=size-2; i>=0; i-- {
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
        res[0] = -1
        stack = append(stack, size-1)

        for i:=1; i<len(nums); i++ {
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
