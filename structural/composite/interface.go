package main

type BaseInstance interface {
	TotalPoints() int
	GetName() string
}

type Team interface {
	AddVictory()
	AddDraw()
	BaseInstance
}

type Order interface {
	OrderByPoints() []BaseInstance
}
