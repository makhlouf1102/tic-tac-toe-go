package main

// CPUPlayer represents an AI player for Ultimate Tic Tac Toe.
type CPUPlayer struct {
	mark Mark
}

// Init initializes the CPU player with its mark.
func (c *CPUPlayer) Init(mark Mark) {
	c.mark = mark
}
