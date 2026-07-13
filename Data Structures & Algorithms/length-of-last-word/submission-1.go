func lengthOfLastWord(s string) int {
	r := []rune(s)

	p2 := len(s) - 1
	p1 := p2
	length := 0

	for p2 >= 0 {
		if r[p2] == ' ' && p1 == p2 {
			p1--
			p2--
			continue
		}

		if r[p2] != ' ' {
			p2--
			length++
			continue
		}

		if r[p2] == ' ' {
			return length
		}
	}
	

	return length
}
