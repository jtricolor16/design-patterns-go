package main

import "fmt"

func main() {
	fmt.Println("state")
	context := NewContext()
	for {
		context.Act()
	}

}
