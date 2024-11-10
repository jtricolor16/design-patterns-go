package main

import "fmt"

type Database interface {
	Save(item string)
	GetAll() []string
}

type Mongo struct{}

func (*Mongo) GetAll() []string {
	return []string{"mongo-a", "mongo-b"}
}

func (*Mongo) Save(item string) {
	fmt.Printf("MONGO - %s\n", item)
}

type Sql struct{}

func (*Sql) GetAll() []string {
	return []string{"sql-a", "sql-b"}
}

func (*Sql) Save(item string) {
	fmt.Printf("SQL - %s\n", item)
}
