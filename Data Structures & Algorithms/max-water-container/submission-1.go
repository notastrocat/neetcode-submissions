func maxArea(heights []int) int {
	maxArea := 0

	left, right := 0, len(heights)-1

	for left < right {
		area := (right - left) * min(heights[left], heights[right])

		if area > maxArea {
			maxArea = area
		}

		if heights[left] < heights[right] {
			left++
		} else if heights[right] < heights[left] {
			right--
		} else {
			left++ 
		}
	}

	return maxArea
}
