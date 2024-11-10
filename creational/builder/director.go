package main

import "sync"

type DirectorDeployment struct {
	builderDeployment BuilderDeployment
}

func NewDirectorDeployment(builderDeployment BuilderDeployment) DirectorDeployment {
	return DirectorDeployment{
		builderDeployment: builderDeployment,
	}
}

func (d DirectorDeployment) ManageBuilder() {
	var wg sync.WaitGroup
	wg.Add(3)
	go d.appConfigRoutine(&wg)
	go d.edgeConfigRoutine(&wg)
	go d.shardConfigRoutine(&wg)
	wg.Wait()
	d.builderDeployment.CreateRingRolloutConfig()
	d.builderDeployment.CreatePlan()
	d.builderDeployment.ExecutePlan()
}

func (d DirectorDeployment) appConfigRoutine(wg *sync.WaitGroup) {
	defer wg.Done()
	d.builderDeployment.CreateAppConfig()
}

func (d DirectorDeployment) edgeConfigRoutine(wg *sync.WaitGroup) {
	defer wg.Done()
	d.builderDeployment.CreateEdgeConfig()
}

func (d DirectorDeployment) shardConfigRoutine(wg *sync.WaitGroup) {
	defer wg.Done()
	d.builderDeployment.CreateShardConfig()
}
