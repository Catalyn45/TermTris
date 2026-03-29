package main

import (
	"time"
)

type Game struct {
	ui    GridUi
	input Input

	currentPiece  *Piece
	grid          [][]int
	score         int
	currentState  State
	startPlayTime time.Time
}

func newGame() *Game {
	grid := make([][]int, config.gameConfig.lines)

	for i := range grid {
		grid[i] = make([]int, config.gameConfig.columns)

		for j := range grid[i] {
			grid[i][j] = 0
		}
	}

	game := &Game{
		input:   newTerminalInput(),
		grid:    grid,
	}

	game.ui = newTerminalGridUi(game)

	return game
}

func (self *Game) start() {
	self.input.initialize()
	self.ui.Initialize()
	self.currentState = newInitialState(self)
	self.startPlayTime = time.Now()

	for {
		now := time.Now().UnixMilli()

		key := self.input.getInput()
		if key == KEY_QUIT {
			return
		}

		self.currentState = self.currentState.Update(key)
		if self.currentState == nil {
			return
		}

		self.ui.Render()

		elapsed := time.Now().UnixMilli() - now
		toWait := int64(1000 / config.gameConfig.ticksPerSecond) - elapsed

		if toWait > 0 {
			time.Sleep(time.Duration(toWait) * time.Millisecond)
		}
	}
}

func (self *Game) UpdateCurrentPiece() {
	for i, row := range self.grid {
		for j, block := range row {
			if block < 0 {
				self.grid[i][j] = 0
			}
		}
	}

	x, y := self.currentPiece.GetPosition()
	shape := self.currentPiece.GetShape()

	xProjection := self.TryGoingDown(shape, x, y)

	for i, row := range shape {
		for j, block := range row {
			if block == 0 {
				continue
			}

			self.grid[xProjection+i][y+j] = PIECE_PROJECTION
			self.grid[x+i][y+j] = -1 * self.currentPiece.pieceType
		}
	}
}

func (self *Game) TryGoingDown(shape [][]int, x int, y int) int {
	xProjection := x
	for i := xProjection; i < len(self.grid); i++ {
		if !self.canMove(shape, i, y) {
			return i - 1
		}
	}

	return len(self.grid) - 1
}

func (self *Game) PlaceCurrentPiece() {
	for i, row := range self.grid {
		for j, block := range row {
			if block < 0 {
				self.grid[i][j] = -1 * block
			}
		}
	}

	self.currentPiece = nil
}

func (self *Game) canMove(shape [][]int, x int, y int) bool {
	for i, row := range shape {
		for j, block := range row {
			if block == 0 {
				continue
			}

			nextI := x + i
			nextJ := y + j

			if nextI < 0 {
				return false
			}

			if nextI >= config.gameConfig.lines {
				return false
			}

			if nextJ < 0 {
				return false
			}

			if nextJ >= config.gameConfig.columns {
				return false
			}

			if self.grid[nextI][nextJ] > 0 {
				return false
			}
		}
	}

	return true
}
