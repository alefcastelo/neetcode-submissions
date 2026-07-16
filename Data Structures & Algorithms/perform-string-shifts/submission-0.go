func stringShift(s string, shift [][]int) string {
	r := []rune(s)

	for i := 0; i < len(shift); i++ {
		direction, amount := shift[i][0], shift[i][1]

		if direction == 0 {

			for j := 0; j < amount; j++ {
				firstElement := r[0]
				r = r[1:]
				r = append(r, firstElement)
			}

			continue
		}


		for j := 0; j < amount; j++ {
			lastElement := r[len(r) -1]
			r = r[:len(r) -1]
			r = append([]rune{lastElement}, r...)
		}
	}

	return string(r)
}
