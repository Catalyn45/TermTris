package main

import "os"
import "golang.org/x/term"
import "bytes"

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
	inputConfig := &config.inputConfig
	select {
		case key := <-self.input:
			if bytes.Equal(key, inputConfig.leftKey) {
				return LEFT
			}

			if bytes.Equal(key, inputConfig.rightKey) {
				return RIGHT
			}

			if bytes.Equal(key, inputConfig.rotateKey) {
				return UP
			}

			if bytes.Equal(key, inputConfig.accelerateKey) {
				return DOWN
			}

			if bytes.Equal(key, inputConfig.quitKey) {
				return Q
			}
		default:
	}

	return NONE
}
