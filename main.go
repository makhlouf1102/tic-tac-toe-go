package main

import (
	"fmt"
	"math/rand"
	"time"
)

const BoardSize = 3 // Standard Tic-Tac-Toe

func main() {
	beginning := time.Now()
	player1 := CPUPlayer{}
	player2 := CPUPlayer{}

	player1.Init(X) // First player is X
	player2.Init(O) // Second player is O

	var p1Wins, p2Wins, draws int

	// Run 1000 games
	for i := 0; i < 10000; i++ {
		board := Board{}
		board.Init(BoardSize)
		winner := PlayGame(&board, &player1, &player2)

		switch winner {
		case X:
			p1Wins++
		case O:
			p2Wins++
		default:
			draws++
		}
	}

	end := time.Since(beginning)

	fmt.Println("Results after 100 games:")
	fmt.Printf("Player 1 (X) Wins: %d\n", p1Wins)
	fmt.Printf("Player 2 (O) Wins: %d\n", p2Wins)
	fmt.Printf("Draws: %d\n", draws)
	fmt.Println(BoardSize)
	fmt.Println(end)

}

func PlayGame(board *Board, p1 *CPUPlayer, p2 *CPUPlayer) Mark {
	currentPlayer := p1

	for !board.IsDone() {
		moves := currentPlayer.GetNextMoveAlphaBeta(*board)

		if len(moves) == 0 {
			break
		}

		move := moves[rand.Intn(len(moves))]
		board.Play(move, currentPlayer.mark)

		if board.isWinner(currentPlayer.mark) {
			return currentPlayer.mark
		}

		if currentPlayer == p1 {
			currentPlayer = p2
		} else {
			currentPlayer = p1
		}
	}

	// fmt.Printf(board.String() + "\n")

	return EMPTY
}
