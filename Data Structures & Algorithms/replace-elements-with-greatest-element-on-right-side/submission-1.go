func replaceElements(arr []int) []int {
	max := len(arr) - 1
	for i := 0; i < max; i++ {
		gt := 0
		
		for j := (i + 1); j <= max; j++ {
			if arr[j] > gt {
				gt = arr[j]
			}
		}

		arr[i] = gt
	}

	arr[max] = -1

	return arr
}