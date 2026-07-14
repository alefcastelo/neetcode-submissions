func validWordSquare(words []string) bool {
	rows, cols := len(words), 0
	for _, row := range words {
		length := len(row)
		
		if length > cols {
			cols = length
		}
	}

	if cols != rows {
		return false
	}

	matrix := make([][]rune, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]rune, cols)
		copy(matrix[i], []rune(words[i]))
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			x := matrix[i][j]
			y := matrix[j][i]

			if x != y {
				return false
			}
		}
	}

	return true
}
