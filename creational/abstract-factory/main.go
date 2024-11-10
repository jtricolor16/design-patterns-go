package main

import "fmt"

func main() {
	executeCreate(NewCommonVehicleFactory())
	executeCreate(NewPremmiumVehicleFactory())

}

func executeCreate(factory AbstractVehicleFactory) {
	fmt.Printf("\n\n")
	car := factory.CreateCar()
	truck := factory.CreateTruck()

	car.Run()
	car.Stop()

	truck.Run()
	truck.Stop()
	truck.TransportItems([]string{"computer", "games", "t-shirts"})
}
