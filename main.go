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

const paddleSymbol = 0x2588
const paddleSize = 4

type Paddle struct {
	row, col, width, heght int
}

var screen tcell.Screen
var Player1 *Paddle
var Player2 *Paddle

func Print(row, col, width, height int, ch rune) {
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			screen.SetContent(col+c, row+r, ch, nil, tcell.StyleDefault)
		}
	}
}

func DrawState() {
	screen.Clear()
	Print(Player1.row, Player1.col, Player1.width, Player1.heght, paddleSymbol)
	Print(Player2.row, Player2.col, Player2.width, Player2.heght, paddleSymbol)
	screen.Show()
}

// This program just prints "Hello, World!".  Press ESC to exit.
func main() {
	InitScreen()
	InitGameState()
	DrawState()

	for {
		switch ev := screen.PollEvent().(type) {
		case *tcell.EventResize:
			screen.Sync()
			DrawState()
		case *tcell.EventKey:
			if ev.Rune() == 'q' {
				screen.Fini()
				os.Exit(0)
			} else if ev.Key() == tcell.KeyUp {
				// TODO: paddle up
			} else if ev.Key() == tcell.KeyDown {
				// TODO: paddle down
			}
		}
	}
}

func InitScreen() {
	encoding.Register()
	var err error
	screen, err = tcell.NewScreen()
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

}

func InitGameState() {
	w, h := screen.Size()
	screen.Clear()
	startPos := h/2 - paddleSize/2

	Player1 = &Paddle{
		row: startPos, col: 0, width: 1, heght: paddleSize,
	}
	Player2 = &Paddle{
		row: startPos, col: w - 1, width: 1, heght: paddleSize,
	}

}
