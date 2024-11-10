package main

import (
	"fmt"
	"time"
)

type VideoGame struct {
	Name                 string
	ReleaseDate          time.Time
	distributionGameType DistributionGameType
}

func (v VideoGame) Play(game string) {
	v.distributionGameType.AddGame(game)
}

func (v VideoGame) Reset() {
	fmt.Printf("Reset the console %s\n", v.Name)
}

func (v VideoGame) TurnOff(game string) {
	v.distributionGameType.SaveGame(game, string(time.Now().GoString()))
	v.distributionGameType.StopGame(game)
	fmt.Printf("turn off the console %s\n", v.Name)
}

/*
super-ninendo - fita
playstation - cd
playstation 2 - dvd
playstation 3 - blueray
playstation 4/5 - blueray + digital
*/
