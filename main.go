package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

// const UltimateBoardSize = 3 // Ultimate board is a 3x3 grid of sub-boards

// func main() {
// 	beginning := time.Now()

// 	var player1, player2 CPUPlayer
// 	player1.Init(X) // CPU player1 uses mark X
// 	player2.Init(O) // CPU player2 uses mark O

// 	var p1Wins, p2Wins, draws int

// 	// Run 1000 games.
// 	for i := 0; i < 1; i++ {
// 		var ub UltimateBoard
// 		ub.Init(UltimateBoardSize)
// 		winner := PlayGame(&ub, &player1, &player2)

// 		switch winner {
// 		case X:
// 			p1Wins++
// 		case O:
// 			p2Wins++
// 		default:
// 			draws++
// 		}
// 	}

// 	elapsed := time.Since(beginning)

// 	fmt.Println("Results after 1000 games:")
// 	fmt.Printf("Player 1 (X) Wins: %d\n", p1Wins)
// 	fmt.Printf("Player 2 (O) Wins: %d\n", p2Wins)
// 	fmt.Printf("Draws: %d\n", draws)
// 	fmt.Println("Ultimate Board Size:", UltimateBoardSize)
// 	fmt.Println("Elapsed time:", elapsed)
// }

// // PlayGame simulates a single Ultimate Tic-Tac-Toe game between two CPU players.
// // It returns the winning mark or EMPTY in case of a draw.
// func PlayGame(ub *UltimateBoard, p1, p2 *CPUPlayer) Mark {
// 	currentPlayer := p1

// 	for !ub.IsDone() {
// 		// Get best moves using alpha-beta search.
// 		moves := currentPlayer.GetNextMoveAlphaBeta(ub)
// 		if len(moves) == 0 {
// 			break
// 		}

// 		// Select one move randomly from the best moves.
// 		move := moves[rand.Intn(len(moves))]
// 		outer := &Move{row: move.OuterRow, col: move.OuterCol}
// 		inner := &Move{row: move.InnerRow, col: move.InnerCol}

// 		ub.Play(outer, inner, currentPlayer.mark)

// 		// Check if current player won on the ultimate board.
// 		if ub.isWinner(currentPlayer.mark) {
// 			return currentPlayer.mark
// 		}

// 		// Switch turns.
// 		if currentPlayer == p1 {
// 			currentPlayer = p2
// 		} else {
// 			currentPlayer = p1
// 		}
// 	}

// 	return EMPTY
// }

// sendCommand sends a command byte followed by an optional message.
func sendCommand(conn net.Conn, cmd byte, message string) error {
	// Send the command byte.
	_, err := conn.Write([]byte{cmd})
	if err != nil {
		return err
	}
	// If there is a message, send it.
	if message != "" {
		_, err = conn.Write([]byte(message))
		if err != nil {
			return err
		}
	}
	return nil
}

// handleConnection handles a single client connection.
func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	// Initialize game state.
	var ub UltimateBoard
	ub.Init(3) // Ultimate board of size 3 (3x3 sub-boards)

	// Assign roles: client is white (X) and server is black (O).
	var clientMark Mark = X
	var serverMark Mark = O
	var cpu CPUPlayer
	cpu.Init(serverMark)

	// Send initial command '1' with the board state.
	boardState := ub.Serialize()
	if err := sendCommand(conn, '1', boardState); err != nil {
		fmt.Println("Error sending initial board:", err)
		return
	}

	// Game loop.
	for {
		// Read client's move (expected format: "r c" with r and c between 0 and 8).
		clientMoveStr, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading client's move:", err)
			return
		}
		clientMoveStr = strings.TrimSpace(clientMoveStr)
		parts := strings.Split(clientMoveStr, " ")
		if len(parts) < 2 {
			sendCommand(conn, '4', "Invalid move format")
			continue
		}
		r, err1 := strconv.Atoi(parts[0])
		c, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil || r < 0 || r >= 9 || c < 0 || c >= 9 {
			sendCommand(conn, '4', "Invalid move coordinates")
			continue
		}

		// Convert overall board coordinates into outer and inner positions.
		outerRow, innerRow := r/3, r%3
		outerCol, innerCol := c/3, c%3

		// Check if the chosen cell is empty.
		if ub.GetAt(outerRow, outerCol, innerRow, innerCol) != EMPTY {
			sendCommand(conn, '4', "Cell not empty")
			continue
		}

		// Make the client's move.
		outer := &Move{row: outerRow, col: outerCol}
		inner := &Move{row: innerRow, col: innerCol}
		ub.Play(outer, inner, clientMark)

		// (In a complete implementation, check for win conditions here.)
		if ub.IsDone() {
			sendCommand(conn, '5', ub.Serialize()+" Client wins")
			break
		}

		// Server (CPU) makes its move.
		cpuMoves := cpu.GetNextMoveAlphaBeta(&ub)
		if len(cpuMoves) == 0 {
			sendCommand(conn, '5', ub.Serialize()+" Draw")
			break
		}
		cpuMove := cpuMoves[0] // For simplicity, choose the first best move.
		sOuter := &Move{row: cpuMove.OuterRow, col: cpuMove.OuterCol}
		sInner := &Move{row: cpuMove.InnerRow, col: cpuMove.InnerCol}
		ub.Play(sOuter, sInner, serverMark)

		if ub.IsDone() {
			sendCommand(conn, '5', ub.Serialize()+" Server wins")
			break
		}

		// Prepare server move info (convert sub-board coordinates back into overall board coordinates).
		moveInfo := fmt.Sprintf("Server move: %d %d", sOuter.row*3+sInner.row, sOuter.col*3+sInner.col)
		// Send command '3' with the updated board state and server's move info.
		if err := sendCommand(conn, '3', ub.Serialize()+" "+moveInfo); err != nil {
			fmt.Println("Error sending server move:", err)
			return
		}
	}
}

// -------------------------
// Main: Start the Server
// -------------------------

func main() {
	ln, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer ln.Close()
	fmt.Println("Server listening on port 8888")
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}
