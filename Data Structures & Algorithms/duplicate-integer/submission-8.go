func hasDuplicate(nums []int) bool {
	sort.Sort(sort.IntSlice(nums))

	for i := 1; i <= len(nums) - 1; i++ {
		if nums[i - 1] == nums[i] {
			return true
		}
	}

	return false
}
