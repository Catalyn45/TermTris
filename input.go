package main

import "os"
import "golang.org/x/term"
import "bytes"

const (
	KEY_NONE = iota
	KEY_ROTATE = iota
	KEY_ACCELERATE = iota
	KEY_LEFT = iota
	KEY_RIGHT = iota
	KEY_INSTANT_DOWN = iota
	KEY_QUIT = iota
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
				return KEY_LEFT
			}

			if bytes.Equal(key, inputConfig.rightKey) {
				return KEY_RIGHT
			}

			if bytes.Equal(key, inputConfig.rotateKey) {
				return KEY_ROTATE
			}

			if bytes.Equal(key, inputConfig.accelerateKey) {
				return KEY_ACCELERATE
			}

			if bytes.Equal(key, inputConfig.instantDownKey) {
				return KEY_INSTANT_DOWN
			}

			if bytes.Equal(key, inputConfig.quitKey) {
				return KEY_QUIT
			}
		default:
	}

	return KEY_NONE
}
