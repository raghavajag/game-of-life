package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
)

type Board struct {
	width  int
	height int
	grid   [][]Cell
}
type Cell struct {
	state bool
}

func (c *Cell) Set(state bool) {
	c.state = state
}

func NewBoard(width, height int) *Board {
	board := Board{
		height: height,
		width:  width,
	}
	board.grid = InitCells(width, height)
	return &board
}
func InitCells(width, height int) [][]Cell {
	cells := make([][]Cell, height)
	for i := 0; i < height; i++ {
		cells[i] = make([]Cell, width)
	}
	return cells
}
func InitScreen() tcell.Screen {
	s, err := tcell.NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating screen: %v\n", err)
		os.Exit(1)
	}
	if err := s.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing scree: %v\n", err)
		os.Exit(1)
	}
	s.Clear()
	s.EnableMouse()
	return s
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
func (b *Board) SetCellState(x, y int, boardGrid [][]bool) {
	for i, row := range boardGrid {
		for j, state := range row {
			if y+i < b.height && x+j < b.width {
				b.grid[y+i][x+j].Set(state)
			}
		}
	}
}
func main() {
	s := InitScreen()
	width, height := s.Size()
	defer s.Fini()
	rand.Seed(time.Now().UnixNano())
	board := NewBoard(width, height)
	board.Random()
	for {
		for i, row := range board.grid {
			for j, cell := range row {
				if cell.state {
					st := tcell.StyleDefault.Background(tcell.ColorWhite)
					s.SetContent(j*2, i, ' ', nil, st)
					s.SetContent(j*2+1, i, ' ', nil, st)
				} else {
					st := tcell.StyleDefault.Background(tcell.ColorBlack)
					s.SetContent(j*2, i, ' ', nil, st)
					s.SetContent(j*2+1, i, ' ', nil, st)
				}
			}
		}
		s.Show()
	}
}
