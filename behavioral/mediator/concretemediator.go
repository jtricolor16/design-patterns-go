package main

import "fmt"

type MatchEvents struct{}

func NewMatchEvents() Mediator {
	return &MatchEvents{}
}

func (m *MatchEvents) Notify(event string, component Component) {
	switch event {
	case "goal":
		component.Show()
	case "offside":
		component.Show()
	case "card":
		component.Show()
	default:
		fmt.Println("invalid event")
	}
}
