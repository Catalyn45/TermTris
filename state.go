package main

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type State interface {
	Update(key int) State
}

func newInitialState(game *Game) State {
	return &InitialState{
		game: game,
	}
}

type InitialState struct {
	game *Game
}

func (self *InitialState) Update(key int) State {
	randomShapeType := rand.Intn(PIECE_COUNT-1) + 1

	piece := newPiece(randomShapeType)
	piece.SetPosition(0, config.gameConfig.columns / 2)
	self.game.currentPiece = piece

	x, y := self.game.currentPiece.GetPosition()
	if !self.game.canMove(self.game.currentPiece.GetShape(), x, y) {
		return nil
	}

	self.game.UpdateCurrentPiece()
	return &PlacingState{
		game: self.game,
	}
}

type PlacingState struct {
	game *Game
	lastMoveTime int64
}

func (self *PlacingState) Update(key int) State {
	now := time.Now().UnixMilli()
	if self.lastMoveTime == 0 {
		self.lastMoveTime = now
	}

	initialX, initialY := self.game.currentPiece.GetPosition()

	newX := initialX
	newY := initialY

	newShape := self.game.currentPiece.GetShape()
	
	accelerate := false
	keyPressed := true

	if key == KEY_LEFT {
		newY -= 1
	} else if key == KEY_RIGHT {
		newY += 1
	} else if key == KEY_ACCELERATE {
		accelerate = true
	} else if key == KEY_INSTANT_DOWN {
		downestX := self.game.TryGoingDown(newShape, newX, newY)
		self.game.currentPiece.SetPosition(downestX, newY)

		self.game.UpdateCurrentPiece()

		initialX, initialY = self.game.currentPiece.GetPosition()
		newX = initialX
		newY = initialY

		accelerate = true
	} else if key == KEY_ROTATE {
		newShape = self.game.currentPiece.Rotate()
	} else {
		keyPressed = false
	}

	// If time passed, try move the piece down
	if  accelerate ||
		now - self.lastMoveTime >= int64(config.statesConfig.timeForMovingOneBlockMilli / config.gameConfig.speed) {
		newX += 1
		self.lastMoveTime = now
	} else if !keyPressed {
		// nothing happened so no reason to continue processing
		return self
	}

	if self.game.canMove(newShape, initialX, initialY) {
		self.game.currentPiece.shape = newShape
	}

	shape := self.game.currentPiece.GetShape()

	canMove := self.game.canMove(shape, newX, newY)
	if !canMove {
		canMove = self.game.canMove(shape, newX, initialY)
		newY = initialY
	}

	if canMove {
		self.game.currentPiece.SetPosition(newX, newY)
		self.game.UpdateCurrentPiece()

		return self
	}

	self.game.PlaceCurrentPiece()
	return &DestroyingState{
		game: self.game,
	}
}

type DestroyingState struct {
	game *Game
}

func (self *DestroyingState) Update(key int) State {
	destroyedLines := []int {}

	for i, row := range self.game.grid {
		lineFull := true
		for _, block := range row {
			if block <= 0 {
				lineFull = false
				break
			}
		}

		if !lineFull {
			continue
		}

		for j := range row {
			self.game.grid[i][j] = 0
		}
		destroyedLines = append(destroyedLines, i)
	}

	if len(destroyedLines) > 0 {
		self.game.score += config.gameConfig.scores[len(destroyedLines) - 1]

		return &FallingState{
			game: self.game,
			destroyedLines: destroyedLines,
		}
	}

	return &InitialState {
		game: self.game,
	}
}

type FallingState struct {
	game *Game
	lastFailingTime int64
	destroyedLines []int
}

func (self *FallingState) Update(key int) State {
	now := time.Now().UnixMilli()
	if self.lastFailingTime == 0 {
		self.lastFailingTime = now
	}

	if now - self.lastFailingTime < int64(config.statesConfig.timeForMovingOneBlockMilli / config.gameConfig.speed) {
		return self;
	}

	firstDestroyedLine := self.destroyedLines[0]
	self.destroyedLines = self.destroyedLines[1:]

	for i := firstDestroyedLine; i > 0; i-- {
		for j := 0; j < config.gameConfig.columns; j++ {
			self.game.grid[i][j] = self.game.grid[i - 1][j]
		}
	}

	for i := 0; i < config.gameConfig.columns; i++ {
		self.game.grid[0][i] = 0
	}

	if len(self.destroyedLines) == 0 {
		return &InitialState{
			game: self.game,
		}
	}

	self.lastFailingTime = now;
	return self
}
