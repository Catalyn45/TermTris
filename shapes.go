package main

const (
	SHAPE_I = iota
	SHAPE_L = iota
	SHAPE_J = iota
	SHAPE_O = iota
	SHAPE_S = iota
	SHAPE_Z = iota
	SHAPE_T = iota
	SHAPE_COUNT = iota
)

var initialShapes = [][][]int {
	SHAPE_I: {
		{1},
		{1},
		{1},
		{1},
	},
	SHAPE_L: {
		{1, 0},
		{1, 0},
		{1, 1},
	},
	SHAPE_J: {
		{0, 1},
		{0, 1},
		{1, 1},
	},
	SHAPE_O: {
		{1, 1},
		{1, 1},
	},
	SHAPE_S: {
		{0, 1, 1},
		{1, 1, 0},
	},
	SHAPE_Z: {
		{1, 1, 0},
		{0, 1, 1},
	},
	SHAPE_T: {
		{1, 1, 1},
		{0, 1, 0},
	},
}

type Shape struct {
	shapeType int

	xPosition int
	yPosition int

	shape [][]int
}

func newShape(shapeType int) *Shape {
	return &Shape{
		shapeType: shapeType,
		shape: initialShapes[shapeType],
	}
}

func (self *Shape) GetPosition() (int, int) {
	return self.xPosition, self.yPosition
}

func (self *Shape) GetShape() [][]int {
	return self.shape
}

func (self *Shape) rotate() {

}
