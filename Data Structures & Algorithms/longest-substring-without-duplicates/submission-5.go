func lengthOfLongestSubstring(s string) int {
	if s=="" {
		return 0
	}
	chars := []rune(s)

	uniquesTable := make(map[rune]int)
	maxLen := 1

	for i:=0; i<len(chars)-1; i++ {
		uniquesTable[chars[i]] = 1
		for j:=i+1; j<len(chars); j++ {
			if _, ok := uniquesTable[chars[j]]; !ok {
				uniquesTable[chars[j]] = 1
				maxLen = max(maxLen, len(uniquesTable))
			} else {
				clear(uniquesTable)
				break
			}
		}
	}

	return maxLen
}
