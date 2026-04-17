func maxArea(heights []int) int {
	maxArea := 0

	for i:=0; i<len(heights); i++ {
		tmpMax := 0
		for j:=i+1; j<len(heights); j++ {
			h := min(heights[i], heights[j])
			area := h * (j-i)

			if area > tmpMax {
				tmpMax = area
			}
		}

		if tmpMax > maxArea {
			maxArea = tmpMax
		}
	}

	return maxArea
}
