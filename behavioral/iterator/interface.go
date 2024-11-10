package main

type GameInterfaceIterator interface {
	HasMore() bool
	NextPhase()
}

type GameInterfaceCollection interface {
	Create() GameInterfaceIterator
}
