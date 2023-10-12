package main

import (
	"github.com/gdamore/tcell/v2"
)

func inputLoop(s tcell.Screen, event chan<- Event) {
	for {
		ev := s.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEnter {
				event <- Event{Type: "done"}
				return
			}
			if ev.Key() == tcell.KeyCtrlSpace {
				event <- Event{Type: "pause"}
				event <- Event{Type: "draw"}
			}
		case *tcell.EventMouse:
			switch ev.Buttons() {
			case tcell.Button1:
				x, y := ev.Position()
				event <- Event{Type: "switchState", X: x / 2, Y: y}
			default:
				continue
			}
		default:
			continue
		}
	}
}
