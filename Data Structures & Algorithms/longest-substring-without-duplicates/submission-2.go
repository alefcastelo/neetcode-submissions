func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	left, right := 0, 1
	longest, r := 1, []rune(s)

	charcounter := make(map[string]int)

	charcounter[string(r[0])] = 1

	for right < len(r) {
		val, exists := charcounter[string(r[right])]

		if exists && val > 0 {
			charcounter[string(r[left])] = charcounter[string(r[left])] - 1
			left++
			continue
		}

		if !exists {
			charcounter[string(r[right])] = 0
		}

		charcounter[string(r[right])] = charcounter[string(r[right])] + 1


		right++

		if right - left > longest {
			longest = right - left
		}
	}

	return longest
}

