package main

type ElementVisitor interface {
	Accept(visitor Visitor)
}

type Visitor interface {
	VisitUser(u *User)
	VisitItem(i *Item)
}
