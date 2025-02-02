package main

type Position interface {
	GetRow() int
	GetCol() int
	SetRow(row int)
	SetCol(col int)
}

type Move struct {
	row int
	col int
}

func (m *Move) GetRow() int {
	return m.row
}

func (m *Move) GetCol() int {
	return m.col
}

func (m *Move) SetRow(row int) {
	m.row = row
}

func (m *Move) SetCol(col int) {
	m.col = col
}
