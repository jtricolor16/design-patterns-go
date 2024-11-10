package main

import "fmt"

func main() {
	fmt.Println("iterator")
	createCollection := NewChairGameCollection(10, 9)
	gameIterator := NewChairGameIterator(createCollection)
	gameIterator.NextPhase()
}
