package main

type GameConfig struct {
	speed int
	tps int
	lines int
	columns int
	scores [4]int
}

type UiConfig struct {
	fps int
	marginDelimiter string
	topdownDelimiter string
	filledBlock string
	emptyBlock string
}

type InputConfig struct {
	leftKey []byte
	rightKey []byte
	rotateKey []byte
	accelerateKey []byte
	instantDownKey []byte
	quitKey []byte
}

type StatesConfig struct {
	timeForMovingOneBlockMilli int
	downAcceleratorMultiplier int
}

type Config struct {
	gameConfig GameConfig
	uiConfig UiConfig
	inputConfig InputConfig
	statesConfig StatesConfig
}

var config = Config {
	gameConfig: GameConfig {
		speed: 1,
		tps: 120,
		lines: 25,
		columns: 10,
		scores: [4]int {100, 300, 500, 800},
	},
	uiConfig: UiConfig {
		fps: 60,
		marginDelimiter: "=",
		topdownDelimiter: "==",
		filledBlock: "[]",
		emptyBlock: "  ",
	},
	inputConfig: InputConfig {
		leftKey: []byte {27, 91, 68}, // Left arrow
		rightKey: []byte {27, 91, 67}, // right arrow
		rotateKey: []byte {27, 91, 65}, // up arrow
		accelerateKey: []byte {27, 91, 66}, // down arrow
		instantDownKey: []byte {' '},
		quitKey: []byte {'q'},
	},
	statesConfig: StatesConfig {
		timeForMovingOneBlockMilli: 1000,
		downAcceleratorMultiplier: 32,
	},
}
