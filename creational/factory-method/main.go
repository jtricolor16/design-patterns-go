package main

import "fmt"

func main() {
	factory := &FactoryMongo{}
	db := factory.CreateDatabase()
	db.Save("test-jeff")
	fmt.Println(db.GetAll())
}
