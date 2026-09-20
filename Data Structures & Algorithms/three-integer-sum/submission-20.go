type IntKey struct {
	A, B, C int
}

func threeSum(nums []int) [][]int {
	sort.Sort(sort.IntSlice(nums))

	hashset := make(map[string]int)

	list := [][]int{}

	for fixed := 0; fixed < len(nums) - 2; fixed++ {

		left, right := fixed+1, len(nums) -1
		
		for left < right {
			a := nums[fixed]
			b := nums[left]
			c := nums[right]

			sum := a + b + c 

			switch {
			case sum == 0:
				key := fmt.Sprint([]int{a, b, c})
				_, exists := hashset[key]
				if !exists {
					list = append(list, []int{a, b, c})
					hashset[key] = 1
				}

				left++
			case sum < 0:
				left++
				continue
			default:
				right--
			}
		}
	}

	return list
}