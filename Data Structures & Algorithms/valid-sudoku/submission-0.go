type comp struct {
    x byte
    y byte
}

func isValidSudoku(board [][]byte) bool {
    rows := make(map[byte]map[byte]bool)
    cols := make(map[byte]map[byte]bool)
    sqrs := make(map[comp]map[byte]bool)

    for r := 0; r < 9; r++ {
        for c := 0; c < 9; c++ {
            val := board[r][c]
            if val == '.' {
                continue
            }

            sCoord := comp{x: byte(r / 3), y: byte(c / 3)}

            if rows[byte(r)] == nil { rows[byte(r)] = make(map[byte]bool) }
            if cols[byte(c)] == nil { cols[byte(c)] = make(map[byte]bool) }
            if sqrs[sCoord] == nil { sqrs[sCoord] = make(map[byte]bool) }

            if rows[byte(r)][val] || cols[byte(c)][val] || sqrs[sCoord][val] {
                return false
            }

            rows[byte(r)][val] = true
            cols[byte(c)][val] = true
            sqrs[sCoord][val] = true
        }
    }

    return true
}