func largestRectangleArea(heights []int) int {
    maxArea := 0

    for i, h := range heights {
        count := 0
        passedI := false // added this to track the "visited" indexes.

        for j := 0; j < len(heights); j++ {
            if heights[j] >= h {
                count++
                if j == i {
                    passedI = true
                }
            } else {
                // Before reaching i, just reset and keep scanning.
                if !passedI {
                    count = 0
                    continue
                }
                // After reaching i, this is the right boundary.
                break
            }
        }

        area := count * h
        if area > maxArea {
            maxArea = area
        }
    }

    return maxArea
}
