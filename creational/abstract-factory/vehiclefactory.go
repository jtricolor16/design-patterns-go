package main

type AbstractVehicleFactory interface {
	CreateCar() Car
	CreateTruck() Truck
}

type PremmiumVehicleFactory struct{}

func NewPremmiumVehicleFactory() AbstractVehicleFactory {
	return &PremmiumVehicleFactory{}
}

func (*PremmiumVehicleFactory) CreateCar() Car {
	return NewFerrari()
}

func (*PremmiumVehicleFactory) CreateTruck() Truck {
	return NewVolvo()
}

type CommonVehicleFactory struct{}

func NewCommonVehicleFactory() AbstractVehicleFactory {
	return &CommonVehicleFactory{}
}

func (*CommonVehicleFactory) CreateCar() Car {
	return NewFiat()
}

func (*CommonVehicleFactory) CreateTruck() Truck {
	return NewScania()
}
