package main

type Board struct {
	matrix [][]Mark
	length int
	winner Mark
}

func (b *Board) Init(size int) {
	b.matrix = make([][]Mark, size)
	for i := 0; i < size; i++ {
		b.matrix[i] = make([]Mark, size)
		for j := 0; j < size; j++ {
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
	length := b.length

	for _, row := range board {
		win := true
		for _, cell := range row {
			if cell != mark {
				win = false
				break
			}
		}
		if win {
			return true
		}
	}

	for colIdx := 0; colIdx < length; colIdx++ {
		win := true
		for rowIdx := 0; rowIdx < length; rowIdx++ {
			if board[rowIdx][colIdx] != mark {
				win = false
				break
			}
		}
		if win {
			return true
		}
	}

	win := true
	for i := 0; i < length; i++ {
		if board[i][i] != mark {
			win = false
			break
		}
	}
	if win {
		return true
	}

	win = true
	for i := 0; i < length; i++ {
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
	positions := []Position{}
	for i, row := range b.matrix {
		for j, cell := range row {
			if cell == EMPTY {
				positions = append(positions, &Move{row: i, col: j})
			}
		}
	}
	return positions
}

func (b *Board) IsDone() bool {
	if b.isWinner(X) || b.isWinner(O) {
		return true
	}
	for _, row := range b.matrix {
		for _, cell := range row {
			if cell == EMPTY {
				return false
			}
		}
	}
	return true
}

func (b *Board) GetWinningPossiblities(mark Mark) int {
	board := b.matrix
	length := b.length
	count := 0

	for row := 0; row < length; row++ {
		possible := true
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
		possible := true
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

	possible := true
	for i := 0; i < length; i++ {
		if board[i][i] != mark {
			possible = false
			break
		}
	}
	if possible {
		count++
	}

	possible = true
	for i := 0; i < length; i++ {
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

func (b *Board) String() string {
	var result string
	for i := 0; i < len(b.matrix); i++ {
		for j := 0; j < len(b.matrix[i]); j++ {
			switch b.matrix[i][j] {
			case X:
				result += "|X"
			case O:
				result += "|O"
			default:
				result += "| "
			}
		}
		result += "|\n"
	}
	return result
}
