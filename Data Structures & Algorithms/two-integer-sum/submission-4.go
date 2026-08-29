func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)

	for i, n := range nums {
		diff := target - n
		if _, exists := hash[diff]; exists {
			return []int{hash[diff], i}
		}

		hash[n] = i
	}

	return []int{}
}
