package drinks

import "decorate/internal/abstract"

type Mocha struct {
	beverage abstract.Beverage
}

func NewMocha(beverage abstract.Beverage) *Mocha {
	return &Mocha{
		beverage: beverage,
	}
}

func (m *Mocha) GetDescription() string {
	return m.beverage.GetDescription() + ", Mocha"
}

func (m *Mocha) Cost() float64 {
	return m.beverage.Cost() + 0.20
}

var _ abstract.CondimentDecorator = (*Mocha)(nil)
