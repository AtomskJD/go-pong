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
	"time"

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
var debugLog string

// This program just prints "Hello, World!".  Press ESC to exit.
func main() {
	InitScreen()
	InitGameState()
	inputChan := InitUserInput()

	cnt := 0
	// MAIN LOOP
	for {
		key := ReadInput(inputChan)
		if key == "Rune[q]" {
			screen.Fini()
			os.Exit(0)
		} else if key == "Rune[w]" {
			Player1.row--
		} else if key == "Rune[s]" {
			Player1.row++
		} else if key == "Up" {
			Player2.row--
		} else if key == "Down" {
			Player2.row++
		}
		DrawState()
		time.Sleep(40 * time.Millisecond)
		cnt++
		debugLog = fmt.Sprintf("%d", cnt)
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

func InitUserInput() chan string {
	inputChan := make(chan string)
	go func() {
		for {
			switch ev := screen.PollEvent().(type) {
			case *tcell.EventKey:
				//debugLog = ev.Name()
				inputChan <- ev.Name()
			}
		}
	}()
	return inputChan
}

func ReadInput(inputChan chan string) string {
	var key string
	select {
	case key = <-inputChan:
	default:
		key = ""
	}

	return key
}

func Print(row, col, width, height int, ch rune) {
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			screen.SetContent(col+c, row+r, ch, nil, tcell.StyleDefault)
		}
	}
}

func PrintScreen(str string) {
	col := 0
	for _, r := range str {
		screen.SetContent(col, 0, r, nil, tcell.StyleDefault)
		col += 1
	}
}

func DrawState() {
	screen.Clear()
	PrintScreen(debugLog)
	Print(Player1.row, Player1.col, Player1.width, Player1.heght, paddleSymbol)
	Print(Player2.row, Player2.col, Player2.width, Player2.heght, paddleSymbol)
	screen.Show()
}
