package abstract

type Beverage interface {
	GetDescription() string
	Cost() float64
}

type CondimentDecorator Beverage
