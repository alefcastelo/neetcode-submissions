func threeSum(nums []int) [][]int {
	sort.Sort(sort.IntSlice(nums))
	hash := make(map[string]int)
	list := [][]int{}

	for i := 0; i < len(nums) - 2; i++ {

		a := nums[i]

		if i >= 1 && nums[i - 1] == a {
			continue
		}		

		l, r := i + 1, len(nums) - 1
		
		for r > l {

			sum := a + nums[l] + nums[r]

			if sum > 0 {
				r--
				continue
			} 
			
			if sum < 0 {
				l++
				continue
			}

			b := nums[l]
			c := nums[r]
			l++

			key := strconv.Itoa(a) + strconv.Itoa(b) + strconv.Itoa(c) 

			_, exists := hash[key]
			
			if exists {
				continue
			}
			
			list = append(list, []int{a, b, c}) 
			hash[key] = 1
		}
	}

	return list
}
