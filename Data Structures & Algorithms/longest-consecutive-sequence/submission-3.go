func longestConsecutive(nums []int) int {

	if len(nums) <  2 {
		return len(nums)
	}

	sort.Ints(nums)
	list := []int{}
	count := 1
	for i := 1; i <= len(nums); i++ {
		if i == len(nums) {
			list = append(list, count)
			break
		}

		if nums[i - 1] + 1 == nums[i] {
			count++
			continue
		}

		if nums[i - 1] == nums[i] {
			continue
		}

		list = append(list, count)
		count = 1
	}

	sort.Sort(sort.Reverse(sort.IntSlice(list)))

	return list[0]
}
