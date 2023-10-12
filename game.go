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
	Screen tcell.Screen
	Board  *Board
	stop   bool
	hide   bool
	Ticker *time.Ticker
	Event  chan Event
	paused bool
	draw   bool
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
		select {
		case ev := <-g.Event:
			switch ev.Type {
			case "switchState":
				if (ev.X < g.Board.width) && (ev.Y < g.Board.height) && (ev.X >= 0) && (ev.Y >= 0) {
					{
						g.Board.Get(ev.X, ev.Y).Switch()
					}
				}
				g.Screen.Show()
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
