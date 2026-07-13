func lengthOfLastWord(s string) int {
	r := []rune(s)

	i, l := len(s) - 1, 0

	for r[i] == ' ' {
		i--
	}

	for i >= 0 && r[i] != ' ' {
		l++
		i--
	}

	return l
}
