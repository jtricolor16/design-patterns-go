package main

import "fmt"

type Requirements struct {
	MinMemory       int
	MinCpuFrequency int
	MinDiskSpace    int
}

type Game struct {
	Name         string
	Requirements Requirements
}

type Console struct {
	Name         string
	Memory       int
	CpuFrequency int
	DiskSpace    int
}

type PS5 struct {
	Console
}

func (p *PS5) Play(game Game) {
	if p.isValidGame(game) {
		fmt.Printf("Play %s on %s\n", game.Name, p.Name)
	} else {
		fmt.Printf("The Console %s does not have specification to run %s", p.Name, game.Name)
	}

}

func (p *PS5) isValidGame(game Game) bool {
	return p.CpuFrequency >= game.Requirements.MinCpuFrequency &&
		p.DiskSpace >= game.Requirements.MinDiskSpace &&
		p.Memory >= game.Requirements.MinMemory
}

func NewPS5() GameStrategyConsole {
	return &PS5{
		Console{
			Name:         "PS5",
			Memory:       32,
			CpuFrequency: 8000,
			DiskSpace:    1000,
		},
	}
}

type CloudGame struct {
	Name         string
	AllowedGames []string
}

func (c *CloudGame) Play(game Game) {
	for _, validGame := range c.AllowedGames {
		if validGame == game.Name {
			fmt.Printf("Play %s on %s\n", game.Name, c.Name)
			return
		}
	}
	fmt.Printf("The Console %s does not have specification to run %s\n", c.Name, game.Name)
}

func NewCloudGame(allowedGames []string) GameStrategyConsole {
	return &CloudGame{
		Name:         "CloudConsole",
		AllowedGames: allowedGames,
	}
}
