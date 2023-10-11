package main

import "math/rand"

type Board struct {
	width  int
	height int
	grid   [][]Cell
	time   int
}

func NewBoard(width, height int) *Board {
	board := Board{
		height: height,
		width:  width,
	}
	board.grid = InitCells(width, height)
	return &board
}

func (b *Board) Random() {
	grid := make([][]bool, b.height)
	for i := 0; i < b.height; i++ {
		grid[i] = make([]bool, b.width)
	}
	for i := 0; i < b.height; i++ {
		for j := 0; j < b.width; j++ {
			grid[i][j] = rand.Int()%2 == 0
		}
	}
	b.SetCellState(0, 0, grid)
}
