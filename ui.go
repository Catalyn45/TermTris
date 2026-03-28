package main

import (
	"fmt"
	"time"
)

type GridUi interface {
	Initialize()
	Render()
}

const topdownDelimiter = "=="
const marginDelimiter = "="

type TerminalGridUi struct {
	cell string
	fps  int
	game *Game

	lastRenderTime int64
}

func newTerminalGridUi(cell string, fps int, game *Game) *TerminalGridUi {
	return &TerminalGridUi{
		cell: cell,
		game: game,
		fps:  fps,
	}
}

func (self *TerminalGridUi) Initialize() {
	fmt.Print("\033[?25l") // hide cursor
}

func (self *TerminalGridUi) Render() {
	nowTime := time.Now().UnixMilli()

	// Check fps
	delta := nowTime - self.lastRenderTime
	frameTime := int64(1000 / self.fps)
	if (delta < frameTime) {
		return
	}

	// Clear screen
    //fmt.Print("\033[2J") // clear screen
    fmt.Print("\033[H")  // move cursor to top-left

	lines := len(self.game.grid)
	columns := len(self.game.grid[0])

	fmt.Print(marginDelimiter)
	for i := 0; i < columns; i++ {
		fmt.Print(topdownDelimiter)
	}
	fmt.Println(marginDelimiter)

	for i := 0; i < lines; i++ {
		fmt.Print(marginDelimiter)
		for j := 0; j < columns; j++ {
			if self.game.grid[i][j] == 0 {
				fmt.Print("  ")
			} else {
				fmt.Print(self.cell)
			}
		}
		fmt.Println(marginDelimiter)
	}

	fmt.Print(marginDelimiter)
	for i := 0; i < columns; i++ {
		fmt.Print(topdownDelimiter)
	}
	fmt.Println(marginDelimiter)

	self.lastRenderTime = nowTime
}
