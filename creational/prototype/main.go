package main

import "fmt"

func main() {
	fmt.Println("Prototype")
	pc := Computer{
		Screen:   "17\"",
		CPU:      "Gabinete com Processador Intel I7",
		Keyboard: "Padrão Brasil",
		Mouse:    "Courser",
		Modem:    "ADSL",
	}
	copy := NewProtypeNotebook(pc)
	copy.Clone()
	fmt.Println(copy.Get())

	copycpu := NewOnlyCPUPrototype(pc)
	copycpu.Clone()
	fmt.Println(copycpu.Get())
}
