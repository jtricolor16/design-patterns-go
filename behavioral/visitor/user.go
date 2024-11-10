package main

import "fmt"

type User struct {
	id   string
	name string
}

func (u *User) Do() string {
	return fmt.Sprintf("%s with id %s\n", u.name, u.id)
}

func (u *User) Accept(v Visitor) {
	v.VisitUser(u)
}

func NewUser(id, name string) *User {
	return &User{
		id:   id,
		name: name,
	}
}
