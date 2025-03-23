package command

import "fmt"

type IStereo interface {
	On()
	Off()
	SetCD()
	SetDVD()
	SetRadio()
	SetVolume(volume int)
}

type Stereo struct {
	CDSlot int
}

func NewStereo() *Stereo {
	return &Stereo{}
}

func (s *Stereo) On() {
	fmt.Println("Stereo is on")
}

func (s *Stereo) Off() {
	fmt.Println("Stereo is off")
}

func (s *Stereo) SetCD() {
	fmt.Println("Stereo is set to CD")
}

func (s *Stereo) SetDVD() {
	fmt.Println("Stereo is set to DVD")
}

func (s *Stereo) SetRadio() {
	fmt.Println("Stereo is set to Radio")
}

func (s *Stereo) SetVolume(volume int) {
	fmt.Println("Stereo is set to volume", volume)
}

type StereoOnWithCDCommand struct {
	Stereo *Stereo
}

func NewStereoOnWithCDCommand(stereo *Stereo) *StereoOnWithCDCommand {
	return &StereoOnWithCDCommand{Stereo: stereo}
}

func (c *StereoOnWithCDCommand) Execute() {
	c.Stereo.On()
	c.Stereo.SetCD()
	c.Stereo.SetVolume(11)
}

type StereoOffCommand struct {
	Stereo *Stereo
}

func NewStereoOffCommand(stereo *Stereo) *StereoOffCommand {
	return &StereoOffCommand{Stereo: stereo}
}

func (c *StereoOffCommand) Execute() {
	c.Stereo.Off()
}
