package main

type Board struct {
	matrix [][]Mark
	length int
}

func (b *Board) Init(size int) {
	b.matrix = make([][]Mark, size)

	for i := 0; i < size; i++ {
		b.matrix[i] = make([]Mark, size)

		for j := range b.matrix[i] {
			b.matrix[i][j] = EMPTY
		}
	}

	b.length = size
}

func (b *Board) Play(p Position, mark Mark) {
	b.matrix[p.GetRow()][p.GetCol()] = mark
}

func (b *Board) UndoPlay(p Position) {
	b.matrix[p.GetRow()][p.GetCol()] = EMPTY
}

func (b *Board) GetAt(row, col int) Mark {
	return b.matrix[row][col]
}

func (b *Board) isWinner(mark Mark) bool {
	board := b.matrix
	win := false
	length := b.length

	for row := range board {
		win = true
		for col := range row {
			if board[row][col] != mark {
				win = false
				break
			}
		}

		if win {
			return true
		}
	}

	for col := range board {
		win = true
		for row := range col {
			if board[row][col] != mark {
				win = false
				break
			}
		}

		if win {
			return true
		}
	}

	win = true
	for i := range board {
		if board[i][i] != mark {
			win = false
			break
		}
	}

	if win {
		return true
	}

	win = true
	for i := range board {
		if board[length-i-1][i] != mark {
			win = false
			break
		}
	}

	if win {
		return true
	}

	return false

}

func (b *Board) Evaluate(mark Mark) int {
	if b.isWinner(mark) {
		return 100
	} else if b.isWinner(GetOpponent(mark)) {
		return -100
	}

	return 0
}

func GetOpponent(mark Mark) Mark {
	if mark == X {
		return O
	}

	return X
}

func (b *Board) GetEmptyCases() []Position {
	positions := make([]Position, 0)

	for row := range b.matrix {
		for col := range row {
			if b.matrix[row][col] == EMPTY {
				positions = append(positions, &Move{row, col})
			}
		}
	}

	return positions
}

func (b *Board) IsDone() bool {
	if b.isWinner(X) || b.isWinner(O) {
		return true
	}

	for row := range b.matrix {
		for col := range row {
			if b.matrix[row][col] == EMPTY {
				return false
			}
		}
	}

	return true
}

func (b *Board) GetWinningPossiblities(mark Mark) int {
	board := b.matrix
	possible := false
	length := b.length
	count := 0

	for row := 0; row < length; row++ {
		possible = true
		for col := 0; col < length; col++ {
			if board[row][col] != mark {
				possible = false
				break
			}
		}

		if possible {
			count++
		}
	}

	for col := 0; col < length; col++ {
		possible = true
		for row := 0; row < length; row++ {
			if board[row][col] != mark {
				possible = false
				break
			}
		}

		if possible {
			count++
		}
	}

	possible = true
	for i := range board {
		if board[i][i] != mark {
			possible = false
			break
		}
	}

	if possible {
		count++
	}

	possible = true
	for i := range board {
		if board[length-i-1][i] != mark {
			possible = false
			break
		}
	}

	if possible {
		count++
	}

	return count
}

func (b *Board) GetMatrix() [][]Mark {
	return b.matrix
}
