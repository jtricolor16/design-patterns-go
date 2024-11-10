package main

import "fmt"

type Car interface {
	Run()
	Stop()
}

type Ferrari struct{}

func NewFerrari() Car {
	return &Ferrari{}
}

func (*Ferrari) Run() {
	fmt.Println("FERRARI: 300km/h")
}

func (*Ferrari) Stop() {
	fmt.Println("FERRARI: 0km/h")
}

type Fiat struct{}

func NewFiat() Car {
	return &Fiat{}
}

func (*Fiat) Run() {
	fmt.Println("FIAT: 140km/h")
}

func (*Fiat) Stop() {
	fmt.Println("FIAT: 0km/h")
}
