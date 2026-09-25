func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers) - 1

	for l < r {
		sum := numbers[l] + numbers[r]

		switch {
		case sum == target:
			return []int{(l+1), (r+1)}
		case sum > target: 
			r--
		default:
			l++
		}
	}

	return []int{}
}
