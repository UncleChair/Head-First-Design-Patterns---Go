package controller

import (
	"command/internal/command"
)

type RemoteControl struct {
	onCommands  []command.Command
	offCommands []command.Command
}

func NewRemoteControl() *RemoteControl {
	return &RemoteControl{
		onCommands:  make([]command.Command, 7),
		offCommands: make([]command.Command, 7),
	}
}

func (r *RemoteControl) SetCommand(slot int, onCommand command.Command, offCommand command.Command) {
	r.onCommands[slot] = onCommand
	r.offCommands[slot] = offCommand
}

func (r *RemoteControl) OnButtonWasPushed(slot int) {
	r.onCommands[slot].Execute()
}

func (r *RemoteControl) OffButtonWasPushed(slot int) {
	r.offCommands[slot].Execute()
}
