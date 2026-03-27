package main

type Game struct {
	speed int

	ui GridUi
	input Input

	grid [][]int

	lines int
	columns int
}

func newGame(lines int, columns int) *Game {
	// create outer slice
	grid := make([][]int, lines)

	// create inner slices
	for i := range grid {
		grid[i] = make([]int, columns)
	}

	return &Game {
		grid: grid,
		ui: newTerminalGridUi(grid, "[]", 20),
		input: newTerminalInput(),
	}
}

func (self *Game) start() {
	self.input.initialize()
	self.ui.Initialize()

	for {

	}
}

func (self *Game) refreshGrid() [][]int {
	return [][]int{{}}
}
