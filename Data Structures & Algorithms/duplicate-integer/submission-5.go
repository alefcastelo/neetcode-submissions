func hasDuplicate(nums []int) bool {
	frequency := map[int]int{}

	for i := 0; i < len(nums); i++ {
		_, exists := frequency[nums[i]]

		if exists {
			return true
		}

		frequency[nums[i]] = 1
	}

	return false
}
