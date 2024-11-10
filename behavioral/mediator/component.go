package main

import "fmt"

type Goal struct {
	id       string
	gameId   string
	player   string
	mediator Mediator
}

func NewGoal(id, gameId string, mediator Mediator) Component {
	return &Goal{
		id:       id,
		gameId:   gameId,
		mediator: mediator,
	}
}

func (g *Goal) GetName() string {
	return "Goal"
}

func (g *Goal) Register(player string) {
	g.player = player
	g.mediator.Notify("goal", g)
}

func (g *Goal) SetPlayer(player string) {
	g.player = player
}

func (g *Goal) Show() {
	fmt.Printf("Player %s score the goal\n", g.player)
}

type Card struct {
	id       string
	gameId   string
	player   string
	color    string
	mediator Mediator
}

func NewCard(id, gameId, color string, mediator Mediator) Component {
	return &Card{
		id:       id,
		gameId:   gameId,
		color:    color,
		mediator: mediator,
	}
}

func (g *Card) GetName() string {
	return "Card"
}

func (g *Card) SetPlayer(player string) {
	g.player = player
}

func (g *Card) Register(player string) {
	g.player = player
	g.mediator.Notify("card", g)
}

func (g *Card) Show() {
	fmt.Printf("Player %s receive the %s card\n", g.player, g.color)
}

type Offside struct {
	id       string
	gameId   string
	player   string
	mediator Mediator
}

func NewOffside(id, gameId string, mediator Mediator) Component {
	return &Offside{
		id:       id,
		gameId:   gameId,
		mediator: mediator,
	}
}

func (g *Offside) GetName() string {
	return "Card"
}

func (g *Offside) Register(player string) {
	g.player = player
	g.mediator.Notify("offside", g)
}

func (g *Offside) SetPlayer(player string) {
	g.player = player
}

func (g *Offside) Show() {
	fmt.Printf("Player %s is offside\n", g.player)
}
