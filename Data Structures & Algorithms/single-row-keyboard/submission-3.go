func calculateTime(keyboard string, word string) int {
	runeK, runeW := []rune(keyboard), []rune(word)
	mapKeyPositions := make(map[string]int)

	for i := 0; i < len(keyboard); i++ {
		mapKeyPositions[string(runeK[i])] = i
	}

	var lastKeyPosition, time int
	for i := 0; i < len(word); i++ {
		charPosition := mapKeyPositions[string(runeW[i])]

		if time == 0 {
			lastKeyPosition = charPosition
			time = charPosition

			continue
		}

		moves := int(math.Abs(float64(lastKeyPosition - charPosition)))
		lastKeyPosition = charPosition
		time = moves + time
	}

	return time
}
