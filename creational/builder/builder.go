package main

import "fmt"

type BuilderDeployment interface {
	CreateAppConfig()
	CreateEdgeConfig()
	CreateShardConfig()
	CreateRingRolloutConfig()
	CreatePlan()
	ExecutePlan()
}

func NewRunRelease() BuilderDeployment {
	return &RunRelease{}
}

type RunRelease struct{}

func (*RunRelease) CreateAppConfig() {
	fmt.Println("Created App Config")
}

func (*RunRelease) CreateEdgeConfig() {
	fmt.Println("Created Edge Config")
}

func (*RunRelease) CreateShardConfig() {
	fmt.Println("Created Shard Config")
}

func (*RunRelease) CreateRingRolloutConfig() {
	fmt.Println("Created RingRollout Config")
}

func (*RunRelease) CreatePlan() {
	fmt.Println("Created Plan")
}

func (*RunRelease) ExecutePlan() {
	fmt.Println("Execute Plan")
}
