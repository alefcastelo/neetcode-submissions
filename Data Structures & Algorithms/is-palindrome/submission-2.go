func isPalindrome(s string) bool {

	s = regexp.MustCompile(`[^a-zA-Z0-9]`).ReplaceAllString(strings.ToLower(s), "")
	r := []rune(s)
	
	i, j := 0, len(s) - 1
	for i < j {
		if r[i] != r[j] {
			return false
		}

		i++
		j--
	}

	return true
}
