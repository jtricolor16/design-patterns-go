package main

import "fmt"

type Factory interface {
	Create(flyweightKind string) DeploymentFlyweight
}

type factoryDeploymentFlyweight struct {
	deploymentFlyweights map[string]DeploymentFlyweight
}

func NewFactoryDeploymentFlyweight() Factory {
	return &factoryDeploymentFlyweight{
		deploymentFlyweights: map[string]DeploymentFlyweight{},
	}
}

func (f *factoryDeploymentFlyweight) Create(flyweightKind string) DeploymentFlyweight {
	flyweightInstance, ok := f.deploymentFlyweights[flyweightKind]
	if !ok {
		switch flyweightKind {
		case "app":
			f.deploymentFlyweights[flyweightKind] = App{definitions: map[string]string{"app": "true"}}
		case "cell":
			f.deploymentFlyweights[flyweightKind] = Cell{definitions: map[string]string{"cell": "true"}}
		case "shard":
			f.deploymentFlyweights[flyweightKind] = Shard{definitions: map[string]string{"shard": "true"}}
		}
		flyweightInstance = f.deploymentFlyweights[flyweightKind]
		fmt.Println("Instanciou ", flyweightKind)
	}
	return flyweightInstance
}
