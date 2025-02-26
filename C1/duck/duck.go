package duck

import "fmt"

// Duck abstract interface
type Duck interface {
	Quack()
	Fly()
}

// FlyBehavior abstract interface
type flyBehavior interface {
	DoFly()
}

// QuackBehavior abstract interface
type quackBehavior interface {
	DoQuack()
}

// GreenDuck type
type GreenDuck struct {
	Color         string
	QuackBehavior quackBehavior
	FlyBehavior   flyBehavior
}

func (d *GreenDuck) Display() {
	fmt.Printf("I am a %s duck🦆\n", d.Color)
}

func (d *GreenDuck) Quack() {
	d.QuackBehavior.DoQuack()
}

func (d *GreenDuck) Fly() {
	d.FlyBehavior.DoFly()
}

// ModelDuck type
type ModelDuck struct {
	Color         string
	QuackBehavior quackBehavior
	FlyBehavior   flyBehavior
}

func (d *ModelDuck) Display() {
	fmt.Printf("I am a %s model duck🦆\n", d.Color)
}

func (d *ModelDuck) SetQuackBehavior(quackBehavior quackBehavior) {
	d.QuackBehavior = quackBehavior
}

func (d *ModelDuck) SetFlyBehavior(flyBehavior flyBehavior) {
	d.FlyBehavior = flyBehavior
}

func (d *ModelDuck) Quack() {
	d.QuackBehavior.DoQuack()
}

func (d *ModelDuck) Fly() {
	d.FlyBehavior.DoFly()
}

// FlyWithWings implements FlyBehavior
type FlyWithWings struct{}

func (f *FlyWithWings) DoFly() {
	fmt.Println("I'm flying with wings!💨")
}

// FlyLikeRocket implements FlyBehavior
type FlyLikeRocket struct{}

func (f *FlyLikeRocket) DoFly() {
	fmt.Println("I'm flying like a rocket!🚀")
}

// FlyNoWay implements FlyBehavior
type FlyNoWay struct{}

func (f *FlyNoWay) DoFly() {
	fmt.Println("I can't fly!")
}

// Quack implements QuackBehavior
type Quack struct{}

func (q *Quack) DoQuack() {
	fmt.Println("Quack!🔈")
}

// Squeak implements QuackBehavior
type Squeak struct{}

func (s *Squeak) DoQuack() {
	fmt.Println("Squeak!🔈")
}

// MuteQuack implements QuackBehavior
type MuteQuack struct{}

func (m *MuteQuack) DoQuack() {
	fmt.Println("<< Silence >>")
}

func DuckImplements() {
	duck := &GreenDuck{
		Color:         "green",
		QuackBehavior: &Quack{},
		FlyBehavior:   &FlyWithWings{},
	}

	// Green duck
	duck.Display()
	duck.Quack()
	duck.Fly()

	modelDuck := &ModelDuck{
		Color:         "black",
		QuackBehavior: &Squeak{},
		FlyBehavior:   &FlyLikeRocket{},
	}

	// Original model duck
	modelDuck.Display()
	modelDuck.Quack()
	modelDuck.Fly()

	fmt.Println("-----Model duck flied for a while-----")

	// Transformed model duck
	modelDuck.Display()
	modelDuck.SetFlyBehavior(&FlyNoWay{})
	modelDuck.SetQuackBehavior(&MuteQuack{})
	// Same as:
	// modelDuck.QuackBehavior = &MuteQuack{}
	// modelDuck.FlyBehavior = &FlyNoWay{}
	modelDuck.Quack()
	modelDuck.Fly()
}
