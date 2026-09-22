func maxArea(heights []int) int {

	a, l, r := 0, 0, len(heights) - 1

	for l < r {
		min := heights[l]

		if heights[r] < heights[l] {
			min = heights[r]
		}

		c := min * (r - l)

		if c > a {
			a = c
		}

		if heights[l] < heights[r] {
			l++
			continue
		}

		r--
	}

	return a
}
