package main

type Mediator interface {
	Notify(event string, component Component)
}

type Component interface {
	GetName() string
	Register(player string)
	Show()
}
