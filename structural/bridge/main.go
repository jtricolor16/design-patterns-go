package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("bridge")
	gameName := "Final Fantasy VII"
	// CD CONSOLE
	cdGame := &CD{}
	ps1 := VideoGame{Name: "Playstation", ReleaseDate: time.Date(1995, time.December, 20, 0, 0, 0, 0, time.Local), distributionGameType: cdGame}
	ps1.Play(gameName)
	ps1.Reset()
	ps1.Play(gameName)
	ps1.TurnOff(gameName)
	fmt.Println("\nDIGITAL CONSOLE\n")
	// DIGITAL CONSOLE
	digital := &Digital{}
	nintendoSwitch := VideoGame{
		Name:                 "Nintendo Switch",
		ReleaseDate:          time.Date(2017, time.April, 0, 0, 0, 0, 0, time.Local),
		distributionGameType: digital,
	}
	nintendoSwitch.Play(gameName)
	nintendoSwitch.Reset()
	nintendoSwitch.TurnOff(gameName)
}
