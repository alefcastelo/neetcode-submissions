func isAnagram(s string, t string) bool {
	i := 0
	max := len(s) - 1
	freq := make(map[string]int)

	if len(s) != len(t) {
		return false
	}

	for i <= max {
		sChar := string(s[i])
		tChar := string(t[i])

		_, sExists := freq[sChar]
		if !sExists {
			freq[sChar] = 1
		} else {
			freq[sChar]++
		}

		_, tExists := freq[tChar]
		if !tExists {
			freq[tChar] = -1
		} else {
			freq[tChar]--
		}

		i++
	}

	for _, val := range freq {
		if val > 0 {
			return false
		}
	}

	return true
}
