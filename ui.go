package main

import "fmt"

type Displayable interface {
	GetPosition() (int, int)
	GetShape() [][]int
}

type GridUi interface {
	Initialize()
	Render()
	AddObject(Displayable)
	RemoveObject(Displayable)
}

type TerminalGridUi struct {
	cell string
	grid [][]int
	fps  int

	lastRenderTime int
	objects        []Displayable
}

func newTerminalGridUi(grid [][]int, cell string, fps int) *TerminalGridUi {
	return &TerminalGridUi{
		cell: cell,
		grid: grid,
		fps:  fps,
	}
}

func (self *TerminalGridUi) Initialize() {

}

func (self *TerminalGridUi) Render() {
	fmt.Print("\033[H\033[2J")
	// Check fps
	// Clear screen
	// Draw outer layer
	// Draw objects
}

func (self *TerminalGridUi) AddObject(object Displayable) {
	self.objects = append(self.objects, object)
}

func (self *TerminalGridUi) RemoveObject(object Displayable) {
	for i, element := range self.objects {
		if element == object { // pointer comparison
			self.objects = append(self.objects[:i], self.objects[i+1:]...)
			return
		}
	}
}
