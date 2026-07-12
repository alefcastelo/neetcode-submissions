func hasDuplicate(nums []int) bool {
	max := len(nums) -1
    first, second := 0, max

	for first < second{
		if nums[first] == nums[second] {
			return true
		}

		second--
		
		if first == second {
			first++
			second = max
		}
	}

	return false
}
