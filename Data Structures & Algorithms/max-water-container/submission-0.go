func maxArea(heights []int) int {
	lastIndex := len(heights)

	currArea := 0

	for i := 0; i < lastIndex; i++ {
		for j := i + 1; j < lastIndex; j++ {
			minHeight := heights[i]
			if heights[j] < heights[i] {
				minHeight = heights[j]
			}

			
			area := (j - i) *  minHeight

			if area < currArea {
				continue
			}

			currArea = area
		}
	} 

	return currArea
}
