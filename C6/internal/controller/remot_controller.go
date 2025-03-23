package controller

import (
	"command/internal/command"
	"fmt"
)

const (
	MaxSlots = 7
)

type RemoteControl struct {
	onCommands  []command.Command
	offCommands []command.Command
	undoCommand command.Command
}

func NewRemoteControl() *RemoteControl {
	return &RemoteControl{
		onCommands:  make([]command.Command, MaxSlots),
		offCommands: make([]command.Command, MaxSlots),
		undoCommand: nil,
	}
}

func (r *RemoteControl) SetCommand(slot int, onCommand command.Command, offCommand command.Command) {
	r.onCommands[slot] = onCommand
	r.offCommands[slot] = offCommand
}

func (r *RemoteControl) OnButtonWasPushed(slot int) {
	if slot < 0 || slot >= MaxSlots {
		fmt.Println("Invalid slot", slot)
		return
	}
	if r.onCommands[slot] != nil {
		r.onCommands[slot].Execute()
		r.undoCommand = r.onCommands[slot]
		return
	}
	fmt.Println("Slot", slot, "is not assigned to any command")
}

func (r *RemoteControl) OffButtonWasPushed(slot int) {
	if slot < 0 || slot >= MaxSlots {
		fmt.Println("Invalid slot", slot)
		return
	}
	if r.offCommands[slot] != nil {
		r.offCommands[slot].Execute()
		r.undoCommand = r.offCommands[slot]
		return
	}
	fmt.Println("Slot", slot, "is not assigned to any command")
}

func (r *RemoteControl) UndoButtonWasPushed() {
	if r.undoCommand != nil {
		r.undoCommand.Undo()
	}
}
