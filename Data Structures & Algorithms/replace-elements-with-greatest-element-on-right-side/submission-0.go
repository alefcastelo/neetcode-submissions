func replaceElements(arr []int) []int {
	i, max := 0, len(arr) - 1
	for i < max {
		r := 0
		j := i + 1
		for j <= max {
			if arr[j] > r {
				r = arr[j]
			}

			j++
		}

		arr[i] = r
		i++
	}

	arr[max] = -1

	return arr
}