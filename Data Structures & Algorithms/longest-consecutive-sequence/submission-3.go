func longestConsecutive(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

	lookup := make(map[int]bool)

	for _, v := range nums {
		lookup[v] = true
	}

	maxLen := 0
	// now i need to iterate over the array and keep checking if the num-1 exists.
	// if it exists, this is NOT the start of the seq.
	for _, v := range nums {
		curLen := 1 // this was the problem in the last submission where
                    // i had to force an increment before returning.
		num := v
		if lookup[v-1] == true {
			continue
		} else {
			num++
			for lookup[num] == true {
				curLen++
				num++
			}
			if curLen > maxLen {
				maxLen = curLen
			}
		}
	}

    // weird workaround but i want to see if the overall logic is somewhat
    // sound, so submitting.
    return maxLen
}