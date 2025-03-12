package main

import (
	"math"
)

// CPUPlayer represents an AI player for Ultimate Tic Tac Toe.
type CPUPlayer struct {
	mark Mark
}

// Init initializes the CPU player with its mark.
func (c *CPUPlayer) Init(mark Mark) {
	c.mark = mark
}

// GetNextMoveAlphaBeta performs an alpha-beta search on the UltimateBoard
// and returns a slice of the best UltimateMove(s) available.
func (c *CPUPlayer) GetNextMoveAlphaBeta(ub *UltimateBoard) []UltimateMove {
	const maxDepth = 3
	moves := make([]UltimateMove, 0)
	emptyMoves := ub.GetEmptyCases()
	bestScore := math.MinInt

	alpha := math.MinInt
	beta := math.MaxInt

	for _, move := range emptyMoves {
		// Convert UltimateMove to outer and inner positions.
		outer := &Move{row: move.OuterRow, col: move.OuterCol}
		inner := &Move{row: move.InnerRow, col: move.InnerCol}

		ub.Play(outer, inner, c.mark)
		score := c.GetScore(ub, maxDepth, false, alpha, beta)
		ub.UndoPlay(outer, inner)

		if score > bestScore {
			moves = []UltimateMove{move}
			bestScore = score
		} else if score == bestScore {
			moves = append(moves, move)
		}
	}

	return moves
}

// GetScore evaluates the ultimate board recursively using alpha-beta pruning.
// The "depth" parameter limits the recursion, and if depth==0 the evaluation is returned.
// isCurrentPlayer indicates whose turn it is in the recursion.
func (c *CPUPlayer) GetScore(ub *UltimateBoard, depth int, isCurrentPlayer bool, alpha, beta int) int {
	if depth == 0 || ub.IsDone() {
		return ub.Evaluate(c.mark)
	}

	emptyMoves := ub.GetEmptyCases()
	var score int
	var bestScore int
	var currentMark Mark

	if isCurrentPlayer {
		bestScore = math.MinInt
		currentMark = c.mark
	} else {
		bestScore = math.MaxInt
		currentMark = GetOpponent(c.mark)
	}

	for _, move := range emptyMoves {
		outer := &Move{row: move.OuterRow, col: move.OuterCol}
		inner := &Move{row: move.InnerRow, col: move.InnerCol}

		ub.Play(outer, inner, currentMark)
		score = c.GetScore(ub, depth-1, !isCurrentPlayer, alpha, beta)
		ub.UndoPlay(outer, inner)

		if isCurrentPlayer {
			bestScore = Max(bestScore, score)
			alpha = Max(alpha, bestScore)
		} else {
			bestScore = Min(bestScore, score)
			beta = Min(beta, bestScore)
		}

		if beta <= alpha {
			break
		}
	}

	return bestScore
}

// Max returns the greater of two integers.
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// Min returns the lesser of two integers.
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
