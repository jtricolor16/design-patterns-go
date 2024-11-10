package main

import "fmt"

func init() {
	fmt.Println("START FLYWEIGHT")
}

func main() {
	factoryInstance := NewFactoryDeploymentFlyweight()
	appFlyweight := factoryInstance.Create("app")
	cellFlyweight := factoryInstance.Create("cell")
	shardFlyweight := factoryInstance.Create("shard")
	context := NewContext([]DeploymentFlyweight{appFlyweight, cellFlyweight, shardFlyweight})
	context.GetDefinitions()
	context.GetResources()

	appFlyweight1 := factoryInstance.Create("app")
	cellFlyweight1 := factoryInstance.Create("cell")
	shardFlyweight1 := factoryInstance.Create("shard")
	context1 := NewContext([]DeploymentFlyweight{appFlyweight1, cellFlyweight1, shardFlyweight1})
	context1.GetDefinitions()
	context1.GetResources()
}
