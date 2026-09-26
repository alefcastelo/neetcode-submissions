func threeSum(nums []int) [][]int {

	list := [][]int{}
	freq := make(map[string]struct{})
	sort.Sort(sort.IntSlice(nums))

	for i := 0; i < len(nums) - 2; i++ {
		j := i + 1
		k := len(nums) - 1

		for j < k {

			a := nums[i]
			b := nums[j]
			c := nums[k]

			sum := a + b + c

			switch {
				case sum == 0:
					item := []int{a, b, c}

					key := fmt.Sprint(item)

					_, exists := freq[key]

					if !exists {
						list = append(list, item)
						freq[key] = struct{}{}
					}

					j++
				case sum < 0:
					j++
				default:
					k--
			}
		}
	}

	return list
}
