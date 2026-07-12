func getConcatenation(nums []int) []int {
    ans := nums
	for _, val := range nums {
		ans = append(ans, val)
	}

	return ans
}
