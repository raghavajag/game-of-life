package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
)

func drawBlock(screen tcell.Screen, x, y int, color tcell.Color) {
	style := tcell.StyleDefault.Foreground(color).Background(tcell.ColorBlack)
	screen.SetContent(x, y, ' ', nil, style)
	screen.SetContent(x+1, y, ' ', nil, style)
	screen.SetContent(x, y+1, ' ', nil, style)
	screen.SetContent(x+1, y+1, ' ', nil, style)
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
func main() {
	s := InitScreen()
	width, height := s.Size()
	defer s.Fini()
	rand.Seed(time.Now().UnixNano())
	for {
		x := rand.Intn(width)
		y := rand.Intn(height)
		color := tcell.Color(rand.Intn(256))
		drawBlock(s, x, y, color)
		s.Show()
		time.Sleep(100 * time.Millisecond)
		// ev := s.PollEvent()
		// switch ev.(type) {
		// case *tcell.EventKey:
		// 	keyEvent := ev.(*tcell.EventKey)
		// 	if keyEvent.Key() == tcell.KeyEnter {
		// 		return
		// 	}
		// default:
		// 	continue
		// }
	}
}
