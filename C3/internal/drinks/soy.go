package drinks

import "decorate/internal/abstract"

type Soy struct {
	beverage abstract.Beverage
}

func NewSoy(beverage abstract.Beverage) *Soy {
	return &Soy{
		beverage: beverage,
	}
}

func (s *Soy) GetDescription() string {
	return s.beverage.GetDescription() + ", Soy"
}

func (s *Soy) Cost() float64 {
	return s.beverage.Cost() + 0.15
}

var _ abstract.CondimentDecorator = (*Soy)(nil)
