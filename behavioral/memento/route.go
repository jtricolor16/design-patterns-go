package main

import "fmt"

type Route struct {
	name string
}

func (r *Route) SetName(name string) {
	r.name = name
}

func (r *Route) Save(name string) MementoInterface {
	return NewRouteMemento(r, name)
}

func NewRoute(name string) *Route {
	return &Route{
		name: name,
	}
}

type RouteMemento struct {
	route *Route
	name  string
}

func (r *RouteMemento) Restore() {
	r.route.SetName(r.name)
}

func (r *RouteMemento) PrintRouteName() {
	fmt.Printf("Current active route: %s\n", r.route.name)
}

func NewRouteMemento(route *Route, oldName string) MementoInterface {
	return &RouteMemento{
		route: route,
		name:  oldName,
	}
}
