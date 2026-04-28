import "slices"

// idek if this is an improvement on the last submission.
func searchMatrix(matrix [][]int, target int) bool {
    for _, row := range matrix {
        first, last := 0, len(row)-1

        if target >= row[first] && target <= row[last] {
            _, found := slices.BinarySearch(row, target)

            if found {
                return true
            }
        }
    }

    return false
}
