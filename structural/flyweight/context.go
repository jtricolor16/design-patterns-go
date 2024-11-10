package main

import "fmt"

type Context struct {
	deploymentComponents []DeploymentFlyweight
}

func NewContext(deploymentComponents []DeploymentFlyweight) Context {
	return Context{
		deploymentComponents: deploymentComponents,
	}
}

func (c Context) GetDefinitions() map[string]string {
	for _, c := range c.deploymentComponents {
		fmt.Println(c.GetDefinitions())
	}
	return map[string]string{}
}
func (c Context) GetResources() map[string]string {
	for _, c := range c.deploymentComponents {
		fmt.Println(c.GetResources())
	}
	return map[string]string{}
}
