package main

import (
	"math/rand"
)

type Board struct {
	width  int
	height int
	grid   [][]Cell
	time   int
}

func (b *Board) SetCellState(x, y int, boardGrid [][]bool) {
	for i, row := range boardGrid {
		for j, state := range row {
			if y+i < b.height && x+j < b.width {
				b.grid[y+i][x+j].Set(state)
			}
		}
	}
}
func (b *Board) Next() {
	for i := 0; i < b.height; i++ {
		for j := 0; j < b.width; j++ {
			b.grid[i][j].CalcNextState()
		}
	}

	for i := 0; i < b.height; i++ {
		for j := 0; j < b.width; j++ {
			b.grid[i][j].Flush()
		}
	}
	b.time++
}
func (b *Board) link() {
	around := [8][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

	// link
	for i := 0; i < b.height; i++ {
		for j := 0; j < b.width; j++ {
			b.grid[i][j].Unlink()
			aroundCells := []*Cell{}
			for _, a := range around {
				y := i + a[0]
				x := j + a[1]

				if x < 0 || y < 0 || x >= b.width || y >= b.height {
					continue
				}

				aroundCells = append(aroundCells, &b.grid[y][x])
			}
			b.grid[i][j].Link(aroundCells)
		}
	}
}
func NewBoard(width, height int) *Board {
	board := Board{
		height: height,
		width:  width,
	}
	board.Init()
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
	b.time = 0
}
func (b *Board) Init() {
	b.grid = Cells(b.width, b.height)
	b.link()
}
func Cells(width, height int) [][]Cell {
	cells := make([][]Cell, height)
	for i := 0; i < height; i++ {
		cells[i] = make([]Cell, width)
	}
	return cells
}
