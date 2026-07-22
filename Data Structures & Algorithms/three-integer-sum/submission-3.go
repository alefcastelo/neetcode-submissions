func threeSum(nums []int) [][]int {
	sort.Sort(sort.IntSlice(nums))
	list := [][]int{}

	hash := make(map[string]int)

	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if i == j || i == k || j == k {
					continue
				}

				if nums[i] + nums[j] + nums[k] == 0 {
					item := []int{nums[i], nums[j], nums[k] }
					itemString := strconv.Itoa(nums[i]) + strconv.Itoa(nums[j]) + strconv.Itoa(nums[i])
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
