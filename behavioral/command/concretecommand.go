package main

type volumeCommand struct {
	value bool
}

func NewCommandVolume(value bool) CommandInterface {
	return &volumeCommand{
		value: value,
	}
}

func (v *volumeCommand) Execute() {
	ChangeVolume(v.value)
}

type turnCommand struct {
	value bool
}

func NewCommandTurn(value bool) CommandInterface {
	return &turnCommand{
		value: value,
	}
}

func (v *turnCommand) Execute() {
	TurnOnTv(v.value)
}
