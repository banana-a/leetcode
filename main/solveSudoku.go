package main

func solveSudoku(board [][]byte) {
	rols := [9][9]bool{}
	cols := [9][9]bool{}
	boxs := [9][9]bool{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				p := board[i][j] - '1'
				rols[i][p] = true
				cols[j][p] = true
				boxs[i/3*3+j/3][p] = true
			}
		}
	}
	var fill func(n int) bool
	fill = func(n int) bool {
		if n == 81 {
			return true
		}
		rol, col := n/9, n%9
		box := (rol/3)*3 + col/3
		if board[rol][col] != '.' {
			return fill(n + 1)
		}
		for i := 0; i < 9; i++ {
			if !rols[rol][i] && !cols[col][i] && !boxs[box][i] {
				rols[rol][i] = true
				cols[col][i] = true
				boxs[box][i] = true
				board[rol][col] = byte(i + '1')
				if fill(n + 1) {
					return true
				}
				rols[rol][i], cols[col][i], boxs[box][i] = false, false, false
			}
		}
		board[rol][col] = '.'
		return false
	}
	fill(0)
}
