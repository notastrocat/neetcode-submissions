// i was trying to write the optimal solution by myself. got stuck at cases where
// there are multiple repeating chars in the string.
// looked at the hint, got even more confused as to how to handle it.
// couple of hours later, this pops into head... why can't we just move past the
// last occurrence? why to loop and remove everything? becuse updating the last
// occurrence of a char in the table should be enough imo.

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	chars := []rune(s)
	lastIndex := make(map[rune]int)

	start := 0
	maxLen := 0

	for end, _ := range chars {
		ch := chars[end]
		li, ok := lastIndex[ch]
		if ok && li >= start {
			// move start just past the previous occurrence
			start = li + 1
		}

		lastIndex[ch] = end
		if cur := end - start + 1; cur > maxLen {
			maxLen = cur
		}
	}

	return maxLen
}
