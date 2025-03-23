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

type LightOnCommand struct {
	Light *Light
}

func NewLightOnCommand(light *Light) *LightOnCommand {
	return &LightOnCommand{
		Light: light,
	}
}

func (l *LightOnCommand) Execute() {
	l.Light.On()
}

type LightOffCommand struct {
	*Light
}

func NewLightOffCommand(light *Light) *LightOffCommand {
	return &LightOffCommand{
		Light: light,
	}
}

func (l *LightOffCommand) Execute() {
	l.Light.Off()
}
