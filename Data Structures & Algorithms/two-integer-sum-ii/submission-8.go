func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]

		switch {
			case target == sum:
				return []int{left + 1, right + 1}
			case target > sum:
				left++
			case target < sum:
				right--
		}
	}

	return []int{left, right}
}