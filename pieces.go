package main

const (
	PIECE_NONE  = iota
	PIECE_I     = iota
	PIECE_L     = iota
	PIECE_J     = iota
	PIECE_O     = iota
	PIECE_S     = iota
	PIECE_Z     = iota
	PIECE_T     = iota
	PIECE_COUNT = iota
)

// special piece for projection
const PIECE_PROJECTION = -1 * PIECE_COUNT

var initialShapes = [][][]int{
	PIECE_I: {
		{0, 1, 0},
		{0, 1, 0},
		{0, 1, 0},
		{0, 1, 0},
	},
	PIECE_L: {
		{0, 1, 0},
		{0, 1, 0},
		{0, 1, 1},
	},
	PIECE_J: {
		{0, 1, 0},
		{0, 1, 0},
		{1, 1, 0},
	},
	PIECE_O: {
		{1, 1},
		{1, 1},
	},
	PIECE_S: {
		{0, 0, 0},
		{0, 1, 1},
		{1, 1, 0},
	},
	PIECE_Z: {
		{1, 1, 0},
		{0, 1, 1},
		{0, 0, 0},
	},
	PIECE_T: {
		{0, 0, 0},
		{1, 1, 1},
		{0, 1, 0},
	},
}

type Piece struct {
	pieceType int

	xPosition int
	yPosition int

	shape [][]int
}

func newPiece(pieceType int) *Piece {
	return &Piece{
		pieceType: pieceType,
		shape:     initialShapes[pieceType],
	}
}

func (self *Piece) GetPosition() (int, int) {
	return self.xPosition, self.yPosition
}

func (self *Piece) SetPosition(x int, y int) {
	self.xPosition = x
	self.yPosition = y
}

func (self *Piece) GetShape() [][]int {
	return self.shape
}

func (self *Piece) Rotate() [][]int {
	lines := len(self.shape)
	columns := len(self.shape[0])

	newShape := make([][]int, columns)
	for i:= range newShape {
		newShape[i] = make([]int, lines)
	}

	for i, row := range self.shape {
		for j, value := range row {
			newShape[columns - j - 1][i] = value
		}
	}

	return newShape
}
