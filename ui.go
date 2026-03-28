package main

import (
	"fmt"
	"time"
)

type GridUi interface {
	Initialize()
	Render()
	DebugMessage(message string)
}

type TerminalGridUi struct {
	game *Game

	lastRenderTime int64

	debugMessage string
}

var helpMessage = []string {
	"Left Arrow  -> Move piece left",
	"Right Arrow -> Move piece right",
	"Down Arrow  -> Piece fall faster",
	"Up Arrow    -> Rotate piece",
	"Space       -> Instant place",
	"q           -> Quit game",
}

func newTerminalGridUi( game *Game) *TerminalGridUi {
	return &TerminalGridUi{
		game: game,
	}
}

func (self *TerminalGridUi) Initialize() {
	fmt.Print("\033[?25l") // hide cursor
    fmt.Print("\033[2J") // clear screen
}

func (self *TerminalGridUi) Render() {
	nowTime := time.Now().UnixMilli()

	// Check fps
	delta := nowTime - self.lastRenderTime
	frameTime := int64(1000 / config.uiConfig.fps)
	if (delta < frameTime) {
		return
	}

	self.Draw()

	self.lastRenderTime = nowTime
}

func (self *TerminalGridUi) DebugMessage(message string) {
	self.debugMessage = message
}

func (self *TerminalGridUi) Draw() {
	// Clear screen
    fmt.Print("\033[H")  // move cursor to top-left

	lines := len(self.game.grid)
	columns := len(self.game.grid[0])

	elapsed := time.Since(self.game.startPlayTime)

	minutes := int(elapsed.Minutes())
	seconds := int(elapsed.Seconds()) % 60

	fmt.Printf("Time: %02d:%02d\n", minutes, seconds)
	fmt.Println("Score: ", self.game.score)

	fmt.Print(config.uiConfig.marginDelimiter)
	for i := 0; i < columns; i++ {
		fmt.Print(config.uiConfig.topdownDelimiter)
	}

	fmt.Println(config.uiConfig.marginDelimiter)

	for i := 0; i < lines; i++ {
		fmt.Print(config.uiConfig.marginDelimiter)
		for j := 0; j < columns; j++ {
			if self.game.grid[i][j] == 0 {
				fmt.Print(config.uiConfig.emptyBlock)
			} else {
				fmt.Print(config.uiConfig.filledBlock)
			}
		}
		fmt.Print(config.uiConfig.marginDelimiter)

		if i < len(helpMessage) {
			fmt.Print(config.uiConfig.emptyBlock, helpMessage[i])
		}

		if self.debugMessage != "" && i == lines - 1 {
			fmt.Print(config.uiConfig.emptyBlock, self.debugMessage)
		}

		fmt.Println("")
	}

	fmt.Print(config.uiConfig.marginDelimiter)
	for i := 0; i < columns; i++ {
		fmt.Print(config.uiConfig.topdownDelimiter)
	}
	fmt.Println(config.uiConfig.marginDelimiter)
}
