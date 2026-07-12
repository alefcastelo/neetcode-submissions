func abs(val int) int {
	if val < 0 {
		return -val
	}

	return val
}

func scoreOfString(s string) int {
	i, j := 0, 1
	b := []byte(s)

	sum := 0
	for i < len(b) && j < len(b) {
		sum = sum + abs(int(b[i]) - int(b[j]))
		i++
		j++
	}

	return sum
}
