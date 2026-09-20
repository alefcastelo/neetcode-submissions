func threeSum(nums []int) [][]int {
	sort.Sort(sort.IntSlice(nums))
	hashset := make(map[string]bool)
	list := [][]int{}

	for pointer := 0; pointer < len(nums) - 2; pointer++ {
		left, right := pointer + 1, len(nums) - 1

		for left < right {
			a := nums[pointer]
			b := nums[left]
			c := nums[right]

			sum := a + b + c

			switch {
				case sum == 0:
					item := []int{a,b,c}
					key := fmt.Sprint(item)
					_, exists := hashset[key]
					if !exists {
						hashset[key] = true
						list = append(list, item)
					}

					left++
				case sum < 0:
					left++
				default:
					right--
			}
		}
	}

	return list
}