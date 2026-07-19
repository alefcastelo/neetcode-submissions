func productExceptSelf(nums []int) []int {
	product, list := 0, []int{}

	for i := 0; i < len(nums); i++ {
		first := true
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}

			if first {
				first = false
				product = nums[j]
				continue
			}

			product *= nums[j]
		}	

		list = append(list, product)
		product = 0
		first = true
	}

	return list

}
