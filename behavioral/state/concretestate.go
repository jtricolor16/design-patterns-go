package main

import (
	"fmt"
	"time"
)

type GreenSemaphore struct {
	context   *Context
	startTime time.Time
}

func (g *GreenSemaphore) Act() {
	currentTime := time.Now()
	then := g.startTime.Add(3 * time.Second)
	if currentTime.Compare(then) > 0 {
		g.context.ChangeState(NewYellowSemaphore(g.context))
	}
	fmt.Println("Run")
}

func NewGreenSemaphore(context *Context) State {
	return &GreenSemaphore{
		context:   context,
		startTime: time.Now(),
	}
}

type YellowSemaphore struct {
	context   *Context
	startTime time.Time
}

func (g *YellowSemaphore) Act() {
	currentTime := time.Now()
	then := g.startTime.Add(2 * time.Second)
	if currentTime.Compare(then) > 0 {
		g.context.ChangeState(NewRedSemaphore(g.context))
	}
	fmt.Println("Attention")
}

func NewYellowSemaphore(context *Context) State {
	return &YellowSemaphore{
		context:   context,
		startTime: time.Now(),
	}
}

type RedSemaphore struct {
	context   *Context
	startTime time.Time
}

func (r *RedSemaphore) Act() {
	currentTime := time.Now()
	then := r.startTime.Add(3 * time.Second)
	if currentTime.Compare(then) > 0 {
		r.context.ChangeState(NewGreenSemaphore(r.context))
	}
	fmt.Println("Stop")
}

func NewRedSemaphore(context *Context) State {
	return &RedSemaphore{
		context:   context,
		startTime: time.Now(),
	}
}
