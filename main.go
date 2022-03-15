// todo:update ball movement
// todo: handle collisions
// todo: handle GAMEOVERS
package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/encoding"
	"os"
	"time"
)

const paddleSymbol = 0x2588
const ballSymbol = 0x25CF
const paddleSize = 4
const initialBallVelocityRow = 1
const initialBallVelocityCol = 2

// merge sturc Ball with Palddle = gameobject
type GameObject struct {
	row, col, width, height int
	velRow, velCol          int
	symbol                  rune
}

var screen tcell.Screen
var Player1 *GameObject
var Player2 *GameObject
var Ball *GameObject

var gameObjects []*GameObject

var debugLog string

// This program just prints "Hello, World!".  Press ESC to exit.
func main() {
	InitScreen()
	InitGameState()
	inputChan := InitUserInput()

	//cnt := 0
	// MAIN LOOP
	for {
		HandleUserInput(ReadInput((inputChan)))

		UpdateState()
		DrawState()

		time.Sleep(80 * time.Millisecond)
		//cnt++
		//debugLog = fmt.Sprintf("%d", cnt)
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

	Player1 = &GameObject{
		row: startPos, col: 0, width: 1, height: paddleSize,
		velRow: 0, velCol: 0,
		symbol: paddleSymbol,
	}
	Player2 = &GameObject{
		row: startPos, col: w - 1, width: 1, height: paddleSize,
		velRow: 0, velCol: 0,
		symbol: paddleSymbol,
	}

	Ball = &GameObject{
		row: h / 2, col: w / 2, height: 1, width: 1,
		velRow: initialBallVelocityRow, velCol: initialBallVelocityCol,
		symbol: ballSymbol,
	}

	gameObjects = []*GameObject{
		Player1, Player2, Ball,
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

func HandleUserInput(key string) {
	_, height := screen.Size()

	if key == "Rune[q]" {
		screen.Fini()
		os.Exit(0)
	} else if key == "Rune[w]" && Player1.row > 0 {
		Player1.row--
	} else if key == "Rune[s]" && Player1.row+Player1.height < height {
		Player1.row++
	} else if key == "Up" && Player2.row > 0 {
		Player2.row--
	} else if key == "Down" && Player2.row+Player2.height < height {
		Player2.row++
	}
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

func UpdateState() {
	for i := range gameObjects {
		gameObjects[i].row += gameObjects[i].velRow
		gameObjects[i].col += gameObjects[i].velCol
	}

	if CollidesWithWall(Ball) {
		Ball.velRow = -Ball.velRow
	}
}

func DrawState() {
	screen.Clear()
	PrintScreen(debugLog)
	for _, obj := range gameObjects {
		Print(obj.row, obj.col, obj.width, obj.height, obj.symbol)
	}
	screen.Show()
}

func CollidesWithWall(obj *GameObject) bool {
	screenWidth, screenHeight := screen.Size()
	_ = screenWidth // dummy
	return obj.row+obj.velRow < 0 || obj.row+obj.velRow >= screenHeight

}
