package main

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
)

const EventTypePlace = "place"

type Event struct {
	Type string
	X    int
	Y    int
}

type Game struct {
	Screen      tcell.Screen
	Board       *Board
	stop        bool
	hide        bool
	Ticker      *time.Ticker
	Event       chan Event
	paused      bool
	draw        bool
	clickedCell *Cell
}

func (g *Game) drawGrid() {
	g.Screen.Clear()

	for y, row := range g.Board.grid {
		for x, alive := range row {
			// Adjust cell coordinates to screen coordinates
			screenX := x * 2
			screenY := y

			if alive.state {
				st := tcell.StyleDefault.Background(tcell.ColorWhite)
				g.Screen.SetContent(screenX, screenY, ' ', nil, st)
				g.Screen.SetContent(screenX+1, screenY, ' ', nil, st)
			} else {
				st := tcell.StyleDefault.Background(tcell.ColorBlack)
				g.Screen.SetContent(screenX, screenY, ' ', nil, st)
				g.Screen.SetContent(screenX+1, screenY, ' ', nil, st)
			}
		}
	}
	g.Screen.Show()
}

func (g *Game) display() {
	g.Screen.Clear()
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
}
func (g *Game) Loop() error {
	for {
		g.display()
		// g.drawGrid()
		select {
		case ev := <-g.Event:
			switch ev.Type {
			case "switchState":
				if (ev.X < g.Board.width) && (ev.Y < g.Board.height) && (ev.X >= 0) && (ev.Y >= 0) {
					cell := g.Board.Get(ev.X, ev.Y)

					// If a cell is already clicked, ignore further clicks until the mouse button is released.
					if g.clickedCell != nil {
						if g.clickedCell == cell {
							// Same cell clicked again, do nothing.
							continue
						}
					}

					// Otherwise, track the new cell and toggle its state.
					g.clickedCell = cell
					cell.Switch()
				}
				g.Screen.Show()
			// case "switchState":
			// 	if (ev.X < g.Board.width) && (ev.Y < g.Board.height) && (ev.X >= 0) && (ev.Y >= 0) {
			// 		{
			// 			g.Board.Get(ev.X, ev.Y).Switch()
			// 		}
			// 	}
			// 	g.Screen.Show()
			case "step":
				if !g.paused {
					g.Board.Next()
				}
				g.Screen.Show()
			case "done":
				return nil
			case "pause":
				g.paused = !g.paused
			case "draw":
				g.draw = !g.draw
			default:
				return fmt.Errorf(ev.Type)
			}
		case <-g.Ticker.C:
			if !g.paused && !g.draw {
				g.Board.Next()
			}
			g.Screen.Show()
		default:
			continue
		}
	}
}
