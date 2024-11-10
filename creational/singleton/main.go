package main

import "fmt"

var singletoneInstace SingletonCounter

func init() {
	singletoneInstace.Count()
}

func main() {
	fmt.Println("singleton")
	singletoneInstace.Count()
	singletoneInstace.Count()
}
