type El struct {
	key int
	val int
}

func topKFrequent(nums []int, k int) []int {
	fm := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		val, exists := fm[nums[i]]
		if exists {
			fm[nums[i]] = val + 1
			continue
		}

		fm[nums[i]] = 1
	}

	list := []El{}

	for key, val := range fm {
		list = append(list, El{
			key: key,
			val: val,
		})
	}

	sort.Slice(list, func (left, right int) bool {
		return list[left].val > list[right].val
	})

	returnable := []int{}
	for i := 0; i < len(list[:k]); i++ {
		key := list[i].key
		returnable = append(returnable, key)
	}

	return returnable
}
