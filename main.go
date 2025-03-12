package main

import (
	"fmt"
	"math/rand"
	"time"
)

const UltimateBoardSize = 3 // Ultimate board is a 3x3 grid of sub-boards

func main() {
	beginning := time.Now()

	var player1, player2 CPUPlayer
	player1.Init(X) // CPU player1 uses mark X
	player2.Init(O) // CPU player2 uses mark O

	var p1Wins, p2Wins, draws int

	// Run 1000 games.
	for i := 0; i < 1000; i++ {
		var ub UltimateBoard
		ub.Init(UltimateBoardSize)
		winner := PlayGame(&ub, &player1, &player2)

		switch winner {
		case X:
			p1Wins++
		case O:
			p2Wins++
		default:
			draws++
		}
	}

	elapsed := time.Since(beginning)

	fmt.Println("Results after 1000 games:")
	fmt.Printf("Player 1 (X) Wins: %d\n", p1Wins)
	fmt.Printf("Player 2 (O) Wins: %d\n", p2Wins)
	fmt.Printf("Draws: %d\n", draws)
	fmt.Println("Ultimate Board Size:", UltimateBoardSize)
	fmt.Println("Elapsed time:", elapsed)
}

// PlayGame simulates a single Ultimate Tic-Tac-Toe game between two CPU players.
// It returns the winning mark or EMPTY in case of a draw.
func PlayGame(ub *UltimateBoard, p1, p2 *CPUPlayer) Mark {
	currentPlayer := p1

	for !ub.IsDone() {
		// Get best moves using alpha-beta search.
		moves := currentPlayer.GetNextMoveAlphaBeta(ub)
		if len(moves) == 0 {
			break
		}

		// Select one move randomly from the best moves.
		move := moves[rand.Intn(len(moves))]
		outer := &Move{row: move.OuterRow, col: move.OuterCol}
		inner := &Move{row: move.InnerRow, col: move.InnerCol}

		ub.Play(outer, inner, currentPlayer.mark)

		// Check if current player won on the ultimate board.
		if ub.isWinner(currentPlayer.mark) {
			return currentPlayer.mark
		}

		// Switch turns.
		if currentPlayer == p1 {
			currentPlayer = p2
		} else {
			currentPlayer = p1
		}
	}

	return EMPTY
}
