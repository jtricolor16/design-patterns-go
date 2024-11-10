package main

func main() {
	mediator := NewMatchEvents()
	firstGoal := NewGoal("1", "1", mediator)
	secondGoal := NewGoal("2", "1", mediator)
	firstCard := NewCard("1", "1", "yellow", mediator)
	firstCard.Register("Jeff")
	firstGoal.Register("Jeff")
	secondGoal.Register("Carol")
}
