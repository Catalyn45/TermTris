package main

import "os"
import "golang.org/x/term"

const (
	NONE = iota
	UP = iota
	DOWN = iota
	LEFT = iota
	RIGHT = iota
	Q = iota
)

type Input interface {
	initialize()
	getInput() int
}

type TerminalInput struct {
	input chan []byte
}

func newTerminalInput() *TerminalInput {
	return &TerminalInput {
		input: make(chan []byte),
	}
}

func (self *TerminalInput) initialize() {
	go func() {
		oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
		if err != nil {
			panic(err)
		}
		defer term.Restore(int(os.Stdin.Fd()), oldState)

		buf := make([]byte, 3)

		for {
			n, _ := os.Stdin.Read(buf)
			self.input <- buf[:n]
		}
	}()
}

func (self *TerminalInput) getInput() int {
	select {
		case key := <-self.input:
			if len(key) == 3 && key[0] == 27 && key[1] == 91 {
				switch key[2] {
					case 65:
						return UP
					case 66:
						return DOWN
					case 67:
						return RIGHT
					case 68:
						return LEFT
				}
			}

			if len(key) == 1 {
				if key[0] == 'q' {
					return Q
				}
			}
		default:
	}

	return NONE
}
