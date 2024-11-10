package main

import "fmt"

type Item struct {
	name, id string
	price    float64
}

func (i *Item) Info() string {
	return fmt.Sprintf("%s - %s: %f\n", i.id, i.name, i.price)
}

func (i *Item) Accept(v Visitor) {
	v.VisitItem(i)
}

func NewItem(name, id string, price float64) *Item {
	return &Item{
		name:  name,
		id:    id,
		price: price,
	}
}
