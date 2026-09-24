// if the window is valid I move the right pointer expanding the window
// the most longest window is the output
// the window formula is the most: ((r - l) - [most freq element]) > k
// if the window is not valid, move the left pointer reducing the window

func maxValueInMap(anyMap map[string]int) int {
	max := 0

	for _, val := range anyMap {
		if val > max {
			max = val
		}
	}

	return max
}

func characterReplacement(s string, k int) int {
	left, right, frequencyMap, runes := 0, 0, make(map[string]int), []rune(s)
	
	longest := 1

	for left <= right && right < len(s) {
		rightChar := string(runes[right])
		_, exists := frequencyMap[rightChar]
		if !exists {
			frequencyMap[rightChar] = 0
		}

		frequencyMap[rightChar] = frequencyMap[rightChar] + 1

		for (right - left) + 1 - maxValueInMap(frequencyMap) > k {
			leftChar := string(runes[left])
			frequencyMap[leftChar] = frequencyMap[leftChar] - 1
			left++
		}

		windowSize := (right - left) + 1

		if windowSize > longest {
			longest = windowSize
		}
		
		right++
	}

	return longest
}
