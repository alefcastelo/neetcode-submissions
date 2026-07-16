func countElements(arr []int) int {
    fm := make(map[int]int)
    
    for i := 0; i < len(arr); i++ {
        _, exists := fm[arr[i]]

        if !exists {
            fm[arr[i]] = 1
        }
    }

    count := 0
    for i := 0; i < len(arr); i++ {
        _, exists := fm[arr[i] + 1]

        if exists {
            count++
        }
        
    }

    return count
}
