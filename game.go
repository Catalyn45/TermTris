package main

type Game struct {
	ui    GridUi
	input Input

	lines   int
	columns int
	speed   int

	currentShape   *Shape
	grid           [][]int
	score          int
	lastUpdateTime int
	currentState   State
}

func newGame(lines int, columns int) *Game {
	// create outer slice
	grid := make([][]int, lines)

	// create inner slices
	for i := range grid {
		grid[i] = make([]int, columns)

		for j := range grid[i] {
			grid[i][j] = 0
		}
	}

	game := &Game{
		input:   newTerminalInput(),
		lines:   lines,
		columns: columns,
		grid:    grid,
		speed:   1,
	}

	game.ui = newTerminalGridUi("[]", 60, game)

	return game
}

func (self *Game) start() {
	self.input.initialize()
	self.ui.Initialize()
	self.currentState = newInitialState(self)

	for {
		key := self.input.getInput()
		if key == Q {
			return
		}

		self.currentState = self.currentState.Update(key)
		if self.currentState == nil {
			return
		}

		self.ui.Render()
	}
}

func (self *Game) UpdateCurrentShape() {
	for i, row := range self.grid {
		for j, block := range row {
			if block < 0 {
				self.grid[i][j] = 0
			}
		}
	}

	x, y := self.currentShape.GetPosition()
	shape := self.currentShape.GetShape()

	for i, row := range shape {
		for j, block := range row {
			if block == 0 {
				continue
			}

			self.grid[x+i][y+j] = -1 * self.currentShape.shapeType
		}
	}
}

func (self *Game) PlaceCurrentShape() {
	for i, row := range self.grid {
		for j, block := range row {
			if block < 0 {
				self.grid[i][j] = -1 * block
			}
		}
	}

	self.currentShape = nil
}

func (self *Game) canMove(x int, y int) bool {
	shape := self.currentShape.GetShape()

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

			if nextI >= self.lines {
				return false
			}

			if nextJ < 0 {
				return false
			}

			if nextJ >= self.columns {
				return false
			}

			if self.grid[nextI][nextJ] > 0 {
				return false
			}
		}
	}

	return true
}

