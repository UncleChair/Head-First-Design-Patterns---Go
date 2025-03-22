package controller

import "command/internal/command"

type SimpleRemoteControl struct {
	slot command.Command
}

func NewSimpleRemoteControl() *SimpleRemoteControl {
	return &SimpleRemoteControl{}
}

func (s *SimpleRemoteControl) SetCommand(command command.Command) {
	s.slot = command
}

func (s *SimpleRemoteControl) ButtonWasPressed() {
	s.slot.Execute()
}
