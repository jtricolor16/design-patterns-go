package main

type Context struct {
	state State
}

func NewContext() *Context {
	context := &Context{}
	context.state = NewGreenSemaphore(context)
	return context
}

func (c *Context) ChangeState(state State) {
	c.state = state
}

func (c *Context) Act() {
	c.state.Act()
}
