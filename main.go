package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/urfave/cli"
)

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
func startGame() error {
	s := InitScreen()
	defer s.Fini()
	width, height := s.Size()
	board := NewBoard(width, height)
	board.Random()
	game := Game{
		Screen: s,
		Board:  board,
	}
	return game.Loop()
}
func main() {
	rand.Seed(time.Now().UnixNano())
	app := cli.NewApp()
	app.Action = func(c *cli.Context) error {
		return startGame()
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}
