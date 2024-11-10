package main

import "fmt"

type Counter struct {
	Value int
}

func NewCounter() *Counter {
	return &Counter{
		Value: 0,
	}
}

type SingletonCounter struct {
	CounterInstance *Counter
}

func (s *SingletonCounter) Count() {
	if s.CounterInstance == nil {
		s.CounterInstance = NewCounter()
	}
	s.CounterInstance.Value++
	fmt.Println(s.CounterInstance.Value)
}
