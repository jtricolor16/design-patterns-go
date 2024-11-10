package main

import "fmt"

func main() {
	fmt.Println("visitor")
	visitorSql := &VisitSql{}
	visitorQueue := &VisitQueue{}
	user := NewUser("1", "jefferson")
	item := NewItem("desodorante", "1", 10.8)
	elements := []ElementVisitor{user, item}
	for _, e := range elements {
		e.Accept(visitorQueue)
		e.Accept(visitorSql)
	}
}
