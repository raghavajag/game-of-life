package main

type Cell struct {
	state bool
}

func (c *Cell) Set(state bool) {
	c.state = state
}

func InitCells(width, height int) [][]Cell {
	cells := make([][]Cell, height)
	for i := 0; i < height; i++ {
		cells[i] = make([]Cell, width)
	}
	return cells
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
