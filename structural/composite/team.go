package main

import "sort"

const (
	victory = 3
	draw    = 1
)

type concreteTeam struct {
	name   string
	points int
}

func NewConcreteTeam(name string) Team {
	return &concreteTeam{
		name:   name,
		points: 0,
	}
}

func (c *concreteTeam) GetName() string {
	return c.name
}

func (c *concreteTeam) AddVictory() {
	c.points += victory
}

func (c *concreteTeam) AddDraw() {
	c.points += draw
}

func (c *concreteTeam) TotalPoints() int {
	return c.points
}

type state struct {
	stateName string
	teams     []Team
}

func NewState(stateName string, teams []Team) BaseInstance {
	return &state{
		stateName: stateName,
		teams:     teams,
	}
}

func (s *state) GetName() string {
	return s.stateName
}

func (s *state) TotalPoints() int {
	var points int
	for _, t := range s.teams {
		points += t.TotalPoints()
	}
	return points
}

type champion struct {
	collection []BaseInstance
}

func NewOrderBy(collection []BaseInstance) Order {
	return &champion{
		collection: collection,
	}
}

func (c *champion) OrderByPoints() []BaseInstance {
	sort.Slice(c.collection, func(i, j int) bool {
		return c.collection[i].TotalPoints() > c.collection[j].TotalPoints()
	})
	return c.collection
}
