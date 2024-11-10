package main

import "fmt"

type Event struct {
	Description string
}

type Subscription interface {
	Notify(event Event)
	GetIdentify() string
}

type ConcreteSubscription struct {
	Name string
}

func (c *ConcreteSubscription) Notify(event Event) {
	fmt.Printf("%s, %s\n", c.Name, event.Description)
}

func (c *ConcreteSubscription) GetIdentify() string {
	return c.Name
}

func NewSubscription(name string) Subscription {
	return &ConcreteSubscription{
		Name: name,
	}
}

type Publisher struct {
	listeners []Subscription
}

func (p *Publisher) AddListener(subscription Subscription) {
	p.listeners = append(p.listeners, subscription)
}

func (p *Publisher) RemoveListener(subscription Subscription) {
	index := -1
	for i, pub := range p.listeners {
		if pub.GetIdentify() == subscription.GetIdentify() {
			index = i
		}
	}
	if index == -1 {
		return
	}
	p.listeners = append(p.listeners[:index], p.listeners[index+1:]...)
}

func (p *Publisher) NotifyAll(event Event) {
	for _, pub := range p.listeners {
		pub.Notify(event)
	}
}
