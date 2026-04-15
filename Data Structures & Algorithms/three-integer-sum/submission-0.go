func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    var res [][]int

    for i := 0; i < len(nums)-2; i++ {
        // Skip duplicate first elements
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }

        seen := make(map[int]bool)
        for j := i + 1; j < len(nums); j++ {
            complement := -nums[i] - nums[j]
            
            if seen[complement] {
                res = append(res, []int{nums[i], complement, nums[j]})
                
                // To avoid duplicate elements, we need to skip
                // consecutive equal elements
                for j+1 < len(nums) && nums[j] == nums[j+1] {
                    j++
                }
            }
            seen[nums[j]] = true
        }
    }
    return res
}