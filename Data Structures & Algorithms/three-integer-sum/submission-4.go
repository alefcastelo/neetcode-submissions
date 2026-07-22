func threeSum(nums []int) [][]int {
	sort.Sort(sort.IntSlice(nums))
	list := [][]int{}

	hash := make(map[string]int)

	for i := 0; i < len(nums)-2; i++ {
		a := nums[i]
		if i >= 1 && nums[i - 1] == a {
			continue
		}

		for j := i + 1; j < len(nums)-1; j++ {
			b := nums[j]
			for k := j + 1; k < len(nums); k++ {
				c := nums[k]
				if i == j || i == k || j == k {
					continue
				}

				if a + b + c == 0 {
					item := []int{a, b, c}
					itemString := strconv.Itoa(a) + strconv.Itoa(b) + strconv.Itoa(c)
					_, exists := hash[itemString]
					if !exists {
						list = append(list, item)
						hash[itemString] = 1 
					}
				}
			}
		}
	}

	return list
}
