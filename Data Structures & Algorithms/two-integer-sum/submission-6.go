func twoSum(nums []int, target int) []int {
    
	freq := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		rest := target - nums[i]

		if key, exists := freq[rest]; exists {
			return []int{key, i}
		}

		freq[nums[i]] = i
	}

	return []int{}
}
