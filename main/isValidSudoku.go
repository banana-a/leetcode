package main

func isValidSudoku(board [][]byte) bool {
	row := make([]map[byte]int, 10)
	col := make([]map[byte]int, 10)
	box := make([]map[byte]int, 10)
	for i := 0; i < 10; i++ {
		row[i] = make(map[byte]int)
		col[i] = make(map[byte]int)
		box[i] = make(map[byte]int)
	}
	for i := 0; i <= 8; i++ {
		for j := 0; j <= 8; j++ {
			row[i][board[i][j]]++
			col[j][board[i][j]]++
			box[(i/3)*3+j/3][board[i][j]]++
			if board[i][j] == '.' {
				continue
			}
			if row[i][board[i][j]] > 1 || col[j][board[i][j]] > 1 || box[(i/3)*3+j/3][board[i][j]] > 1 {
				return false
			}
		}
	}
	return true
}
