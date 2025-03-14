package drinks

import "decorate/internal/abstract"

type Espresso struct {
	description string
	cost        float64
}

func NewEspresso() *Espresso {
	return &Espresso{
		description: "Espresso",
		cost:        1.99,
	}
}

func (e *Espresso) GetDescription() string {
	return e.description
}

func (e *Espresso) Cost() float64 {
	return e.cost
}

var _ abstract.Beverage = (*Espresso)(nil)
