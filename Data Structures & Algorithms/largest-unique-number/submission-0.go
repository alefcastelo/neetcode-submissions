func largestUniqueNumber(nums []int) int {
    freq := make(map[int]int)

    for i := 0; i < len(nums); i++ {
        _, exists := freq[nums[i]]

        if !exists {
            freq[nums[i]] = 1
            continue
        }

        freq[nums[i]] = freq[nums[i]] + 1
    }

    largestNum := -1
    for key, val := range freq {
        if val > 1 {
            continue
        }

        if key > largestNum {
            largestNum = key
        }
    }

    return largestNum
}
