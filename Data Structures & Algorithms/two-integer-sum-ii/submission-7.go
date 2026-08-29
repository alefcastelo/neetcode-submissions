func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		sum := numbers[l] + numbers[r]

		if sum > target {
			r--
			continue
		}

		if sum < target {
			l++
			continue
		}
		
		pl := l + 1
		pr := r + 1

		return []int{pl,pr}
	}

	return []int{}
}