package main

import (
	"encoding/json"
	"os"
)

type GameConfig struct {
	Speed int
	TicksPerSecond int
	Lines int
	Columns int
	Scores [4]int
}

type UiConfig struct {
	Fps int
	MarginDelimiter string
	TopdownDelimiter string
	FilledBlock string
	EmptyBlock string
	ProjectionBlock string
}

type InputConfig struct {
	LeftKey []int
	RightKey []int
	RotateKey []int
	AccelerateKey []int
	InstantDownKey []int
	QuitKey []int
}

type StatesConfig struct {
	TimeForMovingOneBlockMilli int
}

type Config struct {
	GameConfig GameConfig
	UiConfig UiConfig
	InputConfig InputConfig
	StatesConfig StatesConfig
}

var config = Config {
	GameConfig: GameConfig {
		Speed: 1,
		TicksPerSecond: 120,
		Lines: 25,
		Columns: 10,
		Scores: [4]int {100, 300, 500, 800},
	},
	UiConfig: UiConfig {
		Fps: 60,
		MarginDelimiter: "=",
		TopdownDelimiter: "==",
		FilledBlock: "[]",
		EmptyBlock: "  ",
		ProjectionBlock: "{}",
	},
	InputConfig: InputConfig {
		LeftKey: []int {27, 91, 68}, // Left arrow
		RightKey: []int {27, 91, 67}, // right arrow
		RotateKey: []int {27, 91, 65}, // up arrow
		AccelerateKey: []int {27, 91, 66}, // down arrow
		InstantDownKey: []int {' '},
		QuitKey: []int {'q'},
	},
	StatesConfig: StatesConfig {
		TimeForMovingOneBlockMilli: 1000,
	},
}

func updateConfigsFromJson(path string) {
	if path == "" {
		return
	}

    data, err := os.ReadFile(path)
    if err != nil {
		return
    }

    json.Unmarshal(data, &config)
}
