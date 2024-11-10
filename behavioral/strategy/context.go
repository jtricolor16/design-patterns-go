package main

type context struct {
	console GameStrategyConsole
}

func NewContext() *context {
	return &context{}
}

func (c *context) SetConsole(console GameStrategyConsole) {
	c.console = console
}

func (c *context) Run(game Game) {
	c.console.Play(game)
}
