func isAnagram(s string, t string) bool {
	rs := []rune(s)
	rt := []rune(t)

	freq1 := make(map[string]int)
	freq2 := make(map[string]int)

	if len(s) != len(t) {
		return false
	}

	for i := 0; i < len(s); i++ {
		char := string(rs[i])

		_, exists := freq1[char]

		if exists {
			freq1[char] = freq1[char] + 1
		} else {
			freq1[char] = 1
		}
	}

	for i := 0; i < len(t); i++ {
		char := string(rt[i])

		_, exists := freq2[char]

		if exists {
			freq2[char] = freq2[char] + 1
		} else {
			freq2[char] = 1
		}
	}

	for key, val := range freq2 {
		v, exists := freq1[key]

		if !exists {
			return false
		}

		if v != val {
			return false
		}
	}

	return true
}
