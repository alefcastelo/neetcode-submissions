func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	left, right, longest, r, counter := 0, 1, 1, []rune(s), make(map[string]int)

	leftChar := string(r[left])
	rightChar := string(r[right])

	counter[leftChar] = 1

	for right < len(r) {
		leftChar = string(r[left])
		rightChar = string(r[right])

		val, exists := counter[rightChar]
		
		if exists && val > 0 {
			counter[leftChar] = counter[leftChar] - 1
			left++
			continue
		}

		if !exists {
			counter[rightChar] = 0
		}

		counter[rightChar] = counter[rightChar] + 1

		right++

		if right - left > longest {
			longest = right - left
		}
	}

	return longest
}

