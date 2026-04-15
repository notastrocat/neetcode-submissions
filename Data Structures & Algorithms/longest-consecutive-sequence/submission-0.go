func longestConsecutive(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    sort.Ints(nums)

    maxLen := 1
    currentLen := 1

    for i := 1; i < len(nums); i++ {
        if nums[i] == nums[i-1] {
            continue // Skip duplicates
        }

        if nums[i] == nums[i-1]+1 {
            currentLen++
        } else {
            // Sequence broken, record max and reset
            if currentLen > maxLen {
                maxLen = currentLen
            }
            currentLen = 1
        }
    }

    if currentLen > maxLen {
        maxLen = currentLen
    }

    return maxLen
}