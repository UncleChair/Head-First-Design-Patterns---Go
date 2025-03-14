package drinks

import "decorate/internal/abstract"

type Whip struct {
	beverage abstract.Beverage
}

func NewWhip(beverage abstract.Beverage) *Whip {
	return &Whip{
		beverage: beverage,
	}
}

func (w *Whip) GetDescription() string {
	return w.beverage.GetDescription() + ", Whip"
}

func (w *Whip) Cost() float64 {
	return w.beverage.Cost() + 0.10
}

var _ abstract.CondimentDecorator = (*Whip)(nil)
