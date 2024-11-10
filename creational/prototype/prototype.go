package main

type Prototype interface {
	Get() *Computer
	Clone()
}

type Computer struct {
	Screen   string
	Keyboard string
	Mouse    string
	CPU      string
	Modem    string
}

type PrototypeNotebook struct {
	computer Computer
	notebook *Computer
}

func NewProtypeNotebook(computer Computer) Prototype {
	return &PrototypeNotebook{
		computer: computer,
		notebook: new(Computer),
	}
}

func (p *PrototypeNotebook) Clone() {
	p.notebook.Screen = p.computer.Screen
	p.notebook.Keyboard = p.computer.Keyboard
	p.notebook.Modem = p.computer.Modem
}

func (p *PrototypeNotebook) Get() *Computer {
	return p.notebook
}

type OnlyCPUPrototype struct {
	computer Computer
	onlyCpu  *Computer
}

func NewOnlyCPUPrototype(computer Computer) Prototype {
	return &OnlyCPUPrototype{
		computer: computer,
		onlyCpu:  new(Computer),
	}
}

func (o *OnlyCPUPrototype) Clone() {
	o.onlyCpu.CPU = o.computer.CPU
}

func (o *OnlyCPUPrototype) Get() *Computer {
	return o.onlyCpu
}
