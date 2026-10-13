func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)

	for _, val := range nums {
		_, exists := freq[val]

		if !exists {
			freq[val] = 0
		}

		freq[val] = freq[val] + 1
	}

	bucket := make(map[int][]int)

	top := 0
	for key, val := range freq {
		bucket[val] = append(bucket[val], key)

		if val > top {
			top = val
		}
	}

	list := []int{}

	for k > 0 {
		if len(bucket[top]) < 1 {
			top--
			continue
		}

		last := bucket[top][len(bucket[top])-1] 
		bucket[top] = bucket[top][:len(bucket[top])-1]
		list = append(list, last)
		
		k--
	}

	return list
}
