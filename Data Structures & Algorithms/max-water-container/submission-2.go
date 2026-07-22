func maxArea(heights []int) int {
	currArea, l, r := 0, 0, len(heights) - 1
	for l < r {
		minHeight := heights[l]
		if heights[r] < heights[l] {
			minHeight = heights[r]
		}

		area := (r - l) *  minHeight

		if heights[l] < heights[r] {
			l++
		} else {
			r--
		}

		if area < currArea {
			continue
		}

		currArea = area
	} 

	return currArea
}
