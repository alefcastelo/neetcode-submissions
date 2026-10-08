func longestConsecutive(nums []int) int {

	if len(nums) < 1 {
		return 0
	}

	sort.Sort(sort.IntSlice(nums))

	max := 0
	curr := 0
	
	for i := 1; i < len(nums); i++ {

		if nums[i-1] == nums[i] {
			continue
		}

		diff := nums[i] - nums[i - 1]

		if diff > 1 {
			curr = 0
			continue
		}

		curr++

		if curr > max {
			max = curr
		}		
	}

	return max + 1
}
