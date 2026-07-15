import (
	"sort"
)

func largestUniqueNumber(nums []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	
	for i := 0; i < len(nums); i++ {
		if i + 1 == len(nums) {
			return nums[i]
		}

		if nums[i] == nums[i+1] {
			i++
			continue
		}

		return nums[i]
	}

	return -1
}
