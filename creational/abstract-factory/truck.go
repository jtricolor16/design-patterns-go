package main

import "fmt"

type Truck interface {
	Run()
	Stop()
	TransportItems(items []string)
}

type Volvo struct{}

func NewVolvo() Truck {
	return &Volvo{}
}

func (*Volvo) Run() {
	fmt.Println("VOLVO: 200km/h")
}

func (*Volvo) Stop() {
	fmt.Println("VOLVO: 0km/h")
}

func (*Volvo) TransportItems(items []string) {
	fmt.Printf("VOLVO: %v\n", items)
}

type Scania struct{}

func NewScania() Truck {
	return &Scania{}
}

func (*Scania) Run() {
	fmt.Println("SCANIA: 180km/h")
}

func (*Scania) Stop() {
	fmt.Println("SCANIA: 0km/h")
}

func (*Scania) TransportItems(items []string) {
	fmt.Printf("SCANIA: %v\n", items)
}
