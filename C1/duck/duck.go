package duck

import "fmt"

// FlyBehavior abstract interface
type IFlyBehavior interface {
	DoFly()
}

// QuackBehavior abstract interface
type IQuackBehavior interface {
	DoQuack()
}

// Duck abstract interface
type IDuck interface {
	Display()
	Quack()
	Fly()
	SetQuackBehavior(IQuackBehavior)
	SetFlyBehavior(IFlyBehavior)
}

// Duck type
type Duck struct {
	Color string
	// Set behaviors as internal to only allow modification from functions
	quackBehavior IQuackBehavior
	flyBehavior   IFlyBehavior
}

func (d *Duck) Display() {
	fmt.Printf("I am a %s duck🦆\n", d.Color)
}

func (d *Duck) Quack() {
	d.quackBehavior.DoQuack()
}

func (d *Duck) Fly() {
	d.flyBehavior.DoFly()
}

func (d *Duck) SetQuackBehavior(quackBehavior IQuackBehavior) {
	d.quackBehavior = quackBehavior
}

func (d *Duck) SetFlyBehavior(flyBehavior IFlyBehavior) {
	d.flyBehavior = flyBehavior
}

// ModelDuck type
type ModelDuck struct {
	*Duck
}

func NewModelDuck(color string, quackBehavior IQuackBehavior, flyBehavior IFlyBehavior) *ModelDuck {
	return &ModelDuck{
		Duck: &Duck{
			Color:         color,
			quackBehavior: quackBehavior,
			flyBehavior:   flyBehavior,
		},
	}
}

func (d *ModelDuck) Display() {
	fmt.Printf("I am a %s model duck🦆\n", d.Color)
}

// BigDuck type
type BigDuck struct {
	*Duck
}

func NewBigDuck(color string, quackBehavior IQuackBehavior, flyBehavior IFlyBehavior) *BigDuck {
	return &BigDuck{
		Duck: &Duck{
			Color:         color,
			quackBehavior: quackBehavior,
			flyBehavior:   flyBehavior,
		},
	}
}

func (d *BigDuck) Display() {
	fmt.Printf("I am a %s big duck🦆\n", d.Color)
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
	var duck IDuck
	duck = NewBigDuck("green", &Quack{}, &FlyWithWings{})

	// Green big duck
	duck.Display()
	duck.Quack()
	duck.Fly()

	// Change to a new duck
	duck = NewModelDuck("black", &Squeak{}, &FlyLikeRocket{})

	// Original model duck
	duck.Display()
	duck.Quack()
	duck.Fly()

	fmt.Println("-----Model duck flied for a while-----")

	// Transformed model duck
	duck.Display()
	// You could only use these methods to change the behavior
	// of the duck if you are not using a detailed duck type
	duck.SetFlyBehavior(&FlyNoWay{})
	duck.SetQuackBehavior(&MuteQuack{})
	duck.Quack()
	duck.Fly()
}
