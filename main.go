// TODO: Draw paddles
// TODO: Player movements
// TODO:Paddle boundaries
// TODO:Draw ball
// todo:update ball movement
// todo: handle collisions
// todo: handle GAMEOVERS
package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2/encoding"
	"os"

	"github.com/gdamore/tcell/v2"
)

func PrintScreen(screen tcell.Screen, row, col int, str string) {
	for _, c := range str {
		screen.SetContent(col, row, c, nil, tcell.StyleDefault)
		col += 1
	}
}

func Print(screen tcell.Screen, row, col, width, height int, ch rune) {
	for r := 0; r < height; r++ {
		screen.SetContent(col, row+r, ch, nil, tcell.StyleDefault)
	}
}

func displayHelloWorld(screen tcell.Screen) {
	//w, h := screen.Size()
	screen.Clear()
	PrintScreen(screen, 2, 5, "Hello, World!")
	screen.Show()
}

// This program just prints "Hello, World!".  Press ESC to exit.
func main() {
	screen := InitScreen()
	displayHelloWorld(screen)

	for {
		switch ev := screen.PollEvent().(type) {
		case *tcell.EventResize:
			screen.Sync()
			displayHelloWorld(screen)
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyUp {
				screen.Fini()
				os.Exit(0)
			}
		}
	}
}

func InitScreen() tcell.Screen {
	encoding.Register()

	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	if err2 := screen.Init(); err2 != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err2)
		os.Exit(1)
	}

	defStyle := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorMaroon)
	screen.SetStyle(defStyle)

	return screen
}
