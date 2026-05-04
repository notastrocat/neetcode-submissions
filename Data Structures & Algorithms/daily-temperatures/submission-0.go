func dailyTemperatures(temperatures []int) []int {
	maxLen := len(temperatures)
	res := make([]int, maxLen)

	for i:=0; i<maxLen-1; i++ {
		count := 1
		for j:=i+1; j<maxLen; j++ {
			if temperatures[i] < temperatures[j] {
				res[i] = count
				break
			}
			count++
		}
	}

	return res
}

// very naïve way to solve would be to iterate over and over.
// [30,38,30,36,35,40,28]
// [ 0, 1, 2, 3, 4, 5, 6 ]

// [ 1, 4, 1, 2, 1, 0, 0 ]