package main

import (
	"command/internal/command"
	"command/internal/controller"
	"flag"
	"fmt"
)

func main() {
	mode := flag.String("mode", "command", "运行模式: command")
	flag.Parse()

	switch *mode {
	case "command":
		CommandImplements()
	default:
		fmt.Println("Usage: run with -mode=command")
	}
}

func CommandImplements() {
	fmt.Println("This is Command Pattern Implements in Go")
	var (
		remote  = controller.NewSimpleRemoteControl()
		light   = command.NewLight(100)
		lightOn = command.NewLightOnCommand(light)
	)
	remote.SetCommand(lightOn)
	remote.ButtonWasPressed()

	remoteControl := controller.NewRemoteControl()
	livingRoomLight := command.NewLight(100)
	stereo := command.NewStereo()

	livingRoomLightOn := command.NewLightOnCommand(livingRoomLight)
	livingRoomLightOff := command.NewLightOffCommand(livingRoomLight)
	stereoOnWithCD := command.NewStereoOnWithCDCommand(stereo)
	stereoOff := command.NewStereoOffCommand(stereo)

	remoteControl.SetCommand(0, livingRoomLightOn, livingRoomLightOff)
	remoteControl.SetCommand(1, stereoOnWithCD, stereoOff)

	remoteControl.OnButtonWasPushed(0)
	remoteControl.OffButtonWasPushed(0)
	remoteControl.OnButtonWasPushed(1)
	remoteControl.OffButtonWasPushed(1)
}
