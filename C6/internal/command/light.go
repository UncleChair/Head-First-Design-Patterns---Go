package command

import "fmt"

type Light struct {
	Lux int
}

func NewLight(lux int) *Light {
	return &Light{
		Lux: lux,
	}
}

func (l *Light) On() {
	fmt.Println("Light is on")
}

func (l *Light) Off() {
	fmt.Println("Light is off")
}

type LightCommand struct {
	*Light
}

func LightOnCommand(light *Light) *LightCommand {
	return &LightCommand{
		Light: light,
	}
}

func (l *LightCommand) Execute() {
	l.On()
}
