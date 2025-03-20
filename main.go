package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	// writer := bufio.NewWriter(conn)

	for {
		input, err := (reader.ReadByte())
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		// command := strings.TrimSpace(string(input))
		command := input
		switch command {
		case '1':
			fmt.Println("Nouvelle partie! Vous jouez les rouges (X)")
		case '2':
			fmt.Println("Nouvelle partie! Vous jouez les noirs (O)")
		case '3':
			fmt.Println("Coup jouee de la part de l'adversaire ")
		case '4':
			fmt.Println("Coup invalide")
		case '5':
			fmt.Println("Partie Terminee")
		// default:
			// fmt.Println("Commande inconnue :", input)
		}
	}
}

func main() {
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	fmt.Println("Server listening on port 8888")
	defer conn.Close()
	handleConnection(conn)
}
