package main

import (
	"fmt"
	"math/rand"
)

type chairGameIterator struct {
	Collection GameInterfaceCollection
	// could segregate in specific iterator
	People []string
	// could segregate in specific iterator
	Chairs []int
}

func (c *chairGameIterator) HasMore() bool {
	return len(c.People) > 1 && len(c.Chairs) > 0
}

func (c *chairGameIterator) NextPhase() {
	for c.HasMore() {
		valueChair := rand.Intn(len(c.Chairs))
		valuePeople := rand.Intn(len(c.People))
		c.Chairs = append(c.Chairs[:valueChair], c.Chairs[valueChair+1:]...)
		c.People = append(c.People[:valuePeople], c.People[valuePeople+1:]...)
		fmt.Println(c.Chairs, c.People)
	}
	fmt.Printf("The winner is %s\n", c.People[0])
}

func NewChairGameIterator(collection GameInterfaceCollection) GameInterfaceIterator {
	return collection.Create()
}

type chairGameCollection struct {
	numberOfPeople int
	numberOfChairs int
}

func NewChairGameCollection(numberOfPeople, numberOfChairs int) GameInterfaceCollection {
	return &chairGameCollection{
		numberOfPeople: numberOfPeople,
		numberOfChairs: numberOfChairs,
	}
}

func (c *chairGameCollection) Create() GameInterfaceIterator {
	chairGame := &chairGameIterator{
		Collection: c,
		People:     make([]string, c.numberOfPeople),
	}
	for i := 0; i < c.numberOfChairs; i++ {
		chairGame.Chairs = append(chairGame.Chairs, i)
	}
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i := range chairGame.People {
		chairGame.People[i] = string(letterBytes[rand.Intn(len(letterBytes))])
	}
	return chairGame
}
