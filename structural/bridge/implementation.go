package main

import "fmt"

type DistributionGameType interface {
	AddGame(game string)
	StopGame(game string)
	SaveGame(game string, checkpoint string)
}

type CD struct{}

func (*CD) AddGame(game string) {
	fmt.Printf("Insert %s CD on console\n", game)
}

func (*CD) StopGame(game string) {
	fmt.Printf("Remove %s CD from console\n", game)
}

func (*CD) SaveGame(game string, checkpoint string) {
	fmt.Printf("Save game %s at time %s on memory card\n", game, checkpoint)
}

type Digital struct{}

func (*Digital) AddGame(game string) {
	fmt.Printf("Download %s game from internet\n", game)
	fmt.Printf("Select the game %s\n", game)
}

func (*Digital) StopGame(game string) {
	fmt.Printf("Close game %s\n", game)
}

func (*Digital) SaveGame(game string, checkpoint string) {
	fmt.Printf("Save game %s at time %s on console HD\n", game, checkpoint)
}
