package main

import "github.com/gdamore/tcell/v2"

type Game struct {
	Screen tcell.Screen
	Board  *Board
	stop   bool
}

func (g *Game) display() {
	for {
		for i, row := range g.Board.grid {
			for j, cell := range row {
				if cell.state {
					st := tcell.StyleDefault.Background(tcell.ColorWhite)
					g.Screen.SetContent(j*2, i, ' ', nil, st)
					g.Screen.SetContent(j*2+1, i, ' ', nil, st)
				} else {
					st := tcell.StyleDefault.Background(tcell.ColorBlack)
					g.Screen.SetContent(j*2, i, ' ', nil, st)
					g.Screen.SetContent(j*2+1, i, ' ', nil, st)
				}
			}
		}
		g.Screen.Show()
	}
}
func (g *Game) Loop() error {
	for {
		g.display()
	}
}
