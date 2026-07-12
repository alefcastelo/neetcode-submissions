func isSubsequence(s string, t string) bool {
	leftIndex := 0
	rightIndex := 0

	if len(s) == 0 {
		return true
	}

	if  len(t) == 0 {
		return false
	}

	for leftIndex < len(s) - 1 && rightIndex < len(t) - 1 {
		if s[leftIndex] == t[rightIndex] {
			leftIndex++
		}
		
		rightIndex++
	}

	return leftIndex == len(s) - 1
}
