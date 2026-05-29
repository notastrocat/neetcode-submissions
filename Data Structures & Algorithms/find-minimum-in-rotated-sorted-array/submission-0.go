func findMin(nums []int) int {
	// array is just sorted, not rotated. although, the question does say
	// that it'd rotated b/w 1 & n - dunno about the boundaries, so if it's
	// rotated n times... it's the original sorted array. we handle that case here,
	// at the very least.
	if (nums[0] <= nums[len(nums)-1]) {
		return nums[0]
	}

	left := 0
	right := len(nums) - 1

	for left < right {  // we cannot do left <= right this time. find out why?!
        mid := left + (right - left) / 2

        if nums[mid] > nums[right] {
            left = mid + 1
        } else {
            right = mid
        }
    }

    return nums[right]
}
