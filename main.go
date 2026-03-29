package main

import "flag"

func main() {
	userConfigPath := flag.String("config", "", "config file")

    flag.Parse()

	updateConfigsFromJson(*userConfigPath)

	game := newGame()
	game.start()
}