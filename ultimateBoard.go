package main

// UltimateBoard represents the Ultimate Tic Tac Toe board,
// which is a grid of tic-tac-toe boards.
type UltimateBoard struct {
	matrix [][]Board
	length int
}

// Init initializes the ultimate board with the given size.
// Each cell of the ultimate board is itself a 3x3 Board.
func (u *UltimateBoard) Init(size int) {
	u.matrix = make([][]Board, size)
	for i := 0; i < size; i++ {
		u.matrix[i] = make([]Board, size)
		for j := 0; j < size; j++ {
			var board Board
			board.Init(3)
			u.matrix[i][j] = board
		}
	}
	u.length = size
}

// Play makes a move on a sub-board specified by outer (ultimate board)
// and inner (cell within sub-board) positions.
func (u *UltimateBoard) Play(outer Position, inner Position, mark Mark) {
	board := &u.matrix[outer.GetRow()][outer.GetCol()]
	board.Play(inner, mark)
}

// UndoPlay reverts a move on the specified sub-board.
func (u *UltimateBoard) UndoPlay(outer Position, inner Position) {
	board := &u.matrix[outer.GetRow()][outer.GetCol()]
	board.UndoPlay(inner)
}

// GetAt returns the mark at the specified cell within a sub-board.
func (u *UltimateBoard) GetAt(urow, ucol, row, col int) Mark {
	board := u.matrix[urow][ucol]
	return board.GetAt(row, col)
}

// isWinner determines if the ultimate board is won by the given mark.
// It does so by creating a virtual board from sub-board winners.
func (u *UltimateBoard) isWinner(mark Mark) bool {
	length := u.length
	virtual := make([][]Mark, length)
	for i := 0; i < length; i++ {
		virtual[i] = make([]Mark, length)
		for j := 0; j < length; j++ {
			if u.matrix[i][j].isWinner(mark) {
				virtual[i][j] = mark
			} else if u.matrix[i][j].isWinner(GetOpponent(mark)) {
				virtual[i][j] = GetOpponent(mark)
			} else {
				virtual[i][j] = EMPTY
			}
		}
	}

	// Check rows.
	for i := 0; i < length; i++ {
		win := true
		for j := 0; j < length; j++ {
			if virtual[i][j] != mark {
				win = false
				break
			}
		}
		if win {
			return true
		}
	}

	// Check columns.
	for j := 0; j < length; j++ {
		win := true
		for i := 0; i < length; i++ {
			if virtual[i][j] != mark {
				win = false
				break
			}
		}
		if win {
			return true
		}
	}

	// Check main diagonal.
	win := true
	for i := 0; i < length; i++ {
		if virtual[i][i] != mark {
			win = false
			break
		}
	}
	if win {
		return true
	}

	// Check anti-diagonal.
	win = true
	for i := 0; i < length; i++ {
		if virtual[length-i-1][i] != mark {
			win = false
			break
		}
	}
	return win
}

// Evaluate returns a score for the ultimate board state.
func (u *UltimateBoard) Evaluate(mark Mark) int {
	if u.isWinner(mark) {
		return 100
	} else if u.isWinner(GetOpponent(mark)) {
		return -100
	}
	return 0
}

type UltimateMove struct {
	OuterRow int
	OuterCol int
	InnerRow int
	InnerCol int
}

func (u *UltimateBoard) GetEmptyCases() []UltimateMove {
	moves := []UltimateMove{}
	for i := 0; i < u.length; i++ {
		for j := 0; j < u.length; j++ {
			if !u.matrix[i][j].IsDone() {
				empties := u.matrix[i][j].GetEmptyCases() // []Position
				for _, pos := range empties {
					if m, ok := pos.(*Move); ok {
						moves = append(moves, UltimateMove{
							OuterRow: i,
							OuterCol: j,
							InnerRow: m.GetRow(),
							InnerCol: m.GetCol(),
						})
					}
				}
			}
		}
	}
	return moves
}

func (u *UltimateBoard) IsDone() bool {
	if u.isWinner(X) || u.isWinner(O) {
		return true
	}
	for i := 0; i < u.length; i++ {
		for j := 0; j < u.length; j++ {
			if !u.matrix[i][j].IsDone() {
				return false
			}
		}
	}
	return true
}

func (u *UltimateBoard) GetWinningPossiblities(mark Mark) int {
	count := 0
	length := u.length
	virtual := make([][]Mark, length)
	for i := 0; i < length; i++ {
		virtual[i] = make([]Mark, length)
		for j := 0; j < length; j++ {
			if u.matrix[i][j].isWinner(mark) {
				virtual[i][j] = mark
			} else {
				virtual[i][j] = EMPTY
			}
		}
	}

	for i := 0; i < length; i++ {
		possible := true
		for j := 0; j < length; j++ {
			if virtual[i][j] != mark {
				possible = false
				break
			}
		}
		if possible {
			count++
		}
	}

	for j := 0; j < length; j++ {
		possible := true
		for i := 0; i < length; i++ {
			if virtual[i][j] != mark {
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
		if virtual[i][i] != mark {
			possible = false
			break
		}
	}
	if possible {
		count++
	}

	possible = true
	for i := 0; i < length; i++ {
		if virtual[length-i-1][i] != mark {
			possible = false
			break
		}
	}
	if possible {
		count++
	}
	return count
}

func (u *UltimateBoard) GetMatrix() [][]Mark {
	length := u.length
	virtual := make([][]Mark, length)
	for i := 0; i < length; i++ {
		virtual[i] = make([]Mark, length)
		for j := 0; j < length; j++ {
			if u.matrix[i][j].isWinner(X) {
				virtual[i][j] = X
			} else if u.matrix[i][j].isWinner(O) {
				virtual[i][j] = O
			} else {
				virtual[i][j] = EMPTY
			}
		}
	}
	return virtual
}

func (u *UltimateBoard) String() string {
	result := ""
	for i := 0; i < u.length; i++ {
		for subRow := 0; subRow < 3; subRow++ {
			for j := 0; j < u.length; j++ {
				boardStr := u.matrix[i][j].String()
				lines := splitLines(boardStr)
				if subRow < len(lines) {
					result += lines[subRow] + "  "
				}
			}
			result += "\n"
		}
		result += "\n"
	}
	return result
}

func splitLines(s string) []string {
	var lines []string
	current := ""
	for _, c := range s {
		if c == '\n' {
			lines = append(lines, current)
			current = ""
		} else {
			current += string(c)
		}
	}
	if current != "" {
		lines = append(lines, current)
	}
	return lines
}
