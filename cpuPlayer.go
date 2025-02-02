package main

import (
	"math"
)

type CPUPlayer struct {
	mark Mark
}

func (c *CPUPlayer) Init(mark Mark) {
	c.mark = mark
}

func (c *CPUPlayer) GetNextMoveAlphaBeta(board Board) []Position {
	moves := make([]Position, 0)
	emptyMoves := board.GetEmptyCases()
	bestScore := math.MinInt

	alpha := math.MinInt
	beta := math.MaxInt
	for _, move := range emptyMoves {
		board.Play(move, c.mark)
		score := c.GetScore(board, 1, alpha, beta)
		board.UndoPlay(move)

		if score > bestScore {
			moves = []Position{move}
			bestScore = score
		} else if score == bestScore {
			moves = append(moves, move)
		}

	}

	return moves
}

func (c *CPUPlayer) GetScore(board Board, depth, alpha, beta int) int {
	if board.IsDone() {
		// return board.Evaluate(c.mark) * (board.GetWinningPossiblities(GetOpponent(c.mark)) - board.GetWinningPossiblities(c.mark))
		return board.Evaluate(c.mark)
	}

	emptyMoves := board.GetEmptyCases()
	var score int
	var bestScore int

	if depth%2 == 0 {
		bestScore = math.MinInt
		for _, move := range emptyMoves {
			board.Play(move, c.mark)
			score = c.GetScore(board, depth+1, alpha, beta)
			board.UndoPlay(move)
			bestScore = Max(bestScore, score)
			alpha = Max(alpha, bestScore)
			if beta <= alpha {
				break
			}
		}
	} else {
		bestScore = math.MaxInt
		for _, move := range emptyMoves {
			board.Play(move, GetOpponent(c.mark))
			score = c.GetScore(board, depth+1, alpha, beta)
			board.UndoPlay(move)
			bestScore = Min(bestScore, score)
			beta = Min(beta, bestScore)
			if beta <= alpha {
				break
			}
		}
	}

	return bestScore

}

func Max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func Min(x, y int) int {
	if x < y {
		return x
	}

	return y
}
