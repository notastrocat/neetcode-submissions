import "slices"

func searchMatrix(matrix [][]int, target int) bool {
	for _, row := range matrix {
		_, found := slices.BinarySearch(row, target)

		if found {
			return true
		}
	}

	return false
}
