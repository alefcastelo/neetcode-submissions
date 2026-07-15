func anagramMappings(nums1 []int, nums2 []int) []int {
	var mapping []int

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums2[j] == nums1[i] {
				mapping = append(mapping, j)
				break
			}
		} 
	}

	return mapping
}
