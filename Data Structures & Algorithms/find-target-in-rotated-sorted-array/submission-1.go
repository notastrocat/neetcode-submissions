// returns the index of inflection
func getInflection(nums []int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := left + (right - left)/2

		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return right	
}

func bSearch(nums []int, left, right, target int) int {
	for left <= right {
		mid := left + (right - left)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1		
}

// probably a dummy kinda approach to just use what I solved in the last problem.
// might check out the hints to see if there's a better way to solve this one.
func search(nums []int, target int) int {
	inflection := getInflection(nums)
	fmt.Println(inflection)

	found := -1
	found = bSearch(nums, inflection, len(nums)-1, target)

	if -1 == found {
		found = bSearch(nums, 0, inflection, target)
	}

	return found
}
