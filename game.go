package main

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
)

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
	go func() {
		for {
			event := g.Screen.PollEvent()
			switch event.(type) {
			case *tcell.EventKey:
				keyEvent := event.(*tcell.EventKey)
				if keyEvent.Key() == tcell.KeyEnter {
					g.Event <- Event{Type: "done"}
					return
				}
				if keyEvent.Key() == tcell.KeyCtrlSpace {
					g.Event <- Event{Type: "pause"}
				}
			}
		}
	}()

	for {
		g.display()
		select {
		case ev := <-g.Event:
			switch ev.Type {
			case "step":
				if !g.paused {
					g.Board.Next()
				}
				g.Screen.Show()
			case "done":
				return nil
			case "pause":
				g.paused = !g.paused
			default:
				return fmt.Errorf(ev.Type)
			}
		case <-g.Ticker.C:
			if !g.stop && !g.paused {
				g.Board.Next()
			}
			g.Screen.Show()
		}
	}
}
