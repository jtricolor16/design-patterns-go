package main

import "fmt"

type VisitSql struct{}

func (v *VisitSql) VisitUser(u *User) {
	fmt.Printf("Save on sql User table: %s\n", u.Do())
}

func (v *VisitSql) VisitItem(i *Item) {
	fmt.Printf("Save on sql Item table: %s\n", i.Info())
}

type VisitQueue struct{}

func (v *VisitQueue) VisitUser(u *User) {
	fmt.Printf("Publish on User queue: %s\n", u.Do())
}

func (v *VisitQueue) VisitItem(i *Item) {
	fmt.Printf("Publish on Item queue: %s\n", i.Info())
}
