func hasDuplicate(nums []int) bool {
	sort.Sort(sort.IntSlice(nums))

	i := 1
	for i < len(nums) {
		if nums[i-1] == nums[i] {
			return true
		}

		i++
	}

	return false
}
