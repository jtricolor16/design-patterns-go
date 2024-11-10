package main

import "fmt"

func main() {
	fmt.Println("strategy")
	game1 := Game{
		Name: "Persona 5 Royal",
		Requirements: Requirements{
			MinMemory:       8,
			MinCpuFrequency: 800,
			MinDiskSpace:    100,
		},
	}
	game2 := Game{
		Name: "FF7 Rebirth",
		Requirements: Requirements{
			MinMemory:       16,
			MinCpuFrequency: 6000,
			MinDiskSpace:    500,
		},
	}
	context := NewContext()
	cloudGame := NewCloudGame([]string{"ForzaHorizon5", "Persona 5 Royal"})
	PS5 := NewPS5()

	// check on cloud game
	context.SetConsole(cloudGame)
	context.Run(game1)
	context.Run(game2)

	// check on ps5
	context.SetConsole(PS5)
	context.Run(game1)
	context.Run(game2)
}
