package main

import "fmt"

func main() {
	fmt.Println("template-method")
	pc := NewPC("Personal Computer Samsung")
	cellPhone := NewCellPhone("Galaxy S23")
	ConnectToNetwork(pc)
	ConnectToNetwork(cellPhone)
}
