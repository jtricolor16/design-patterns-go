package main

import (
	"fmt"
)

func init() {
	fmt.Println("START COMPOSITE")
}

func main() {
	fmt.Println("\nCAMPEONATO BRASILEIRO")

	fluminense := NewConcreteTeam("Fluminense")
	flamengo := NewConcreteTeam("Flamengo")
	vasco := NewConcreteTeam("Vasco")
	botafogo := NewConcreteTeam("Botafogo")
	rioDeJaneiro := NewState("Rio de Janeiro", []Team{fluminense, flamengo, vasco, botafogo})

	saoPaulo := NewConcreteTeam("São Paulo")
	palmeiras := NewConcreteTeam("Palmeiras")
	corinthians := NewConcreteTeam("Corinthians")
	santos := NewConcreteTeam("Santos")
	saoPauloState := NewState("São Paulo", []Team{saoPaulo, palmeiras, corinthians, santos})

	fluminense.AddVictory()
	fluminense.AddVictory()
	fluminense.AddVictory()
	fluminense.AddVictory()
	fluminense.AddDraw()

	saoPaulo.AddVictory()
	saoPaulo.AddVictory()
	saoPaulo.AddVictory()
	saoPaulo.AddDraw()
	saoPaulo.AddDraw()

	botafogo.AddVictory()
	botafogo.AddDraw()
	botafogo.AddDraw()
	botafogo.AddDraw()
	botafogo.AddDraw()

	santos.AddDraw()
	santos.AddDraw()
	santos.AddVictory()
	santos.AddVictory()

	flamengo.AddDraw()
	flamengo.AddDraw()
	flamengo.AddDraw()

	corinthians.AddVictory()

	palmeiras.AddVictory()
	palmeiras.AddVictory()
	palmeiras.AddDraw()

	vasco.AddDraw()

	teams := []BaseInstance{flamengo, fluminense, botafogo, vasco, palmeiras, saoPaulo, santos, corinthians}

	fmt.Printf("\nCHAMPIONS BY TEAM\n")

	championsTeam := NewOrderBy(teams)

	for i, champs := range championsTeam.OrderByPoints() {
		fmt.Println(i+1, champs.GetName(), champs.TotalPoints())
	}

	fmt.Printf("\nCHAMPIONS BY STATE\n")

	championsState := NewOrderBy([]BaseInstance{rioDeJaneiro, saoPauloState})

	for i, champs := range championsState.OrderByPoints() {
		fmt.Println(i+1, champs.GetName(), champs.TotalPoints())
	}
}
