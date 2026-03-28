package main

import (
	"math/rand"
	"time"
)

func init() {
	// rand.Seed(time.Now().Unix())
	rand.Seed(1)
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
	randomShapeType := rand.Intn(SHAPE_COUNT-1) + 1

	shape := newShape(randomShapeType)
	shape.SetPosition(0, self.game.columns / 2)
	self.game.currentShape = shape

	x, y := self.game.currentShape.GetPosition()
	if !self.game.canMove(x, y) {
		return nil
	}

	return &PlacingState{
		game: self.game,
	}
}

type PlacingState struct {
	game *Game
	lastMoveTime int64
}

func (self *PlacingState) Update(key int) State {
	initialX, initialY := self.game.currentShape.GetPosition()
	now := time.Now().UnixMilli()

	newX := initialX
	newY := initialY

	// If time passed, try move the piece down
	if self.lastMoveTime == 0 || now - self.lastMoveTime >= int64(1000 / self.game.speed ) {
		newX += 1
		self.lastMoveTime = now
	}

	speed := 1
	if key == LEFT {
		newY -= 1
	} else if key == RIGHT {
		newY += 1
	} else if key == DOWN {
		speed = 16
	}

	self.game.speed = speed

	canMove := self.game.canMove(newX, newY)
	if !canMove {
		canMove = self.game.canMove(newX, initialY)
		newY = initialY
	}

	if canMove {
		self.game.currentShape.SetPosition(newX, newY)
		self.game.UpdateCurrentShape()

		return self
	}

	self.game.PlaceCurrentShape()
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

	if now - self.lastFailingTime < int64(1000 / self.game.speed) {
		return self;
	}

	firstDestroyedLine := self.destroyedLines[0]
	self.destroyedLines = self.destroyedLines[1:]

	for i := firstDestroyedLine; i > 0; i-- {
		for j := 0; j < self.game.columns; j++ {
			self.game.grid[i][j] = self.game.grid[i - 1][j]
		}
	}

	for i := 0; i < self.game.columns; i++ {
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
