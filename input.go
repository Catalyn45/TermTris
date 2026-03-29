package main

import "os"
import "golang.org/x/term"

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

func equalByteInt(b []byte, ints []int) bool {
    if len(b) != len(ints) {
        return false
    }

    for i := range b {
        if int(b[i]) != ints[i] {
            return false
        }
    }

    return true
}

func (self *TerminalInput) getInput() int {
	inputConfig := &config.InputConfig
	select {
		case key := <-self.input:
			if equalByteInt(key, inputConfig.LeftKey) {
				return KEY_LEFT
			}

			if equalByteInt(key, inputConfig.RightKey) {
				return KEY_RIGHT
			}

			if equalByteInt(key, inputConfig.RotateKey) {
				return KEY_ROTATE
			}

			if equalByteInt(key, inputConfig.AccelerateKey) {
				return KEY_ACCELERATE
			}

			if equalByteInt(key, inputConfig.InstantDownKey) {
				return KEY_INSTANT_DOWN
			}

			if equalByteInt(key, inputConfig.QuitKey) {
				return KEY_QUIT
			}
		default:
	}

	return KEY_NONE
}
