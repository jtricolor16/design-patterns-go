package main

type App struct {
	resources   map[string]string
	definitions map[string]string
}

func (a App) GetDefinitions() map[string]string {
	return a.definitions
}

func (a App) GetResources() map[string]string {
	return a.resources
}

type Shard struct {
	resources   map[string]string
	definitions map[string]string
}

func (a Shard) GetDefinitions() map[string]string {
	return a.definitions
}

func (a Shard) GetResources() map[string]string {
	return a.resources
}

type Cell struct {
	resources   map[string]string
	definitions map[string]string
}

func (a Cell) GetDefinitions() map[string]string {
	return a.definitions
}

func (a Cell) GetResources() map[string]string {
	return a.resources
}
