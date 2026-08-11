func isValid(s string) bool {

	if len(s) < 2 {
		return false
	}

	popMapping := map[string]string{
		"]": "[",
		")": "(",
		"}": "{",
	}

	stack := []string{}

	r := []rune(s)

	for _, valRune := range r {
		val := string(valRune)

		popChar, exists := popMapping[val]

		if !exists {
			stack = append(stack, val)
			continue
		}

		if len(stack) == 0 {
			return false
		}

		lastChar := stack[len(stack)-1] 
		
		if popChar != lastChar {
			return false
		}

		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}
