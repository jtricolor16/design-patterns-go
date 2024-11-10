package main

type DeploymentFlyweight interface {
	GetDefinitions() map[string]string
	GetResources() map[string]string
}
