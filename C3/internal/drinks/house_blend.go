package drinks

import "decorate/internal/abstract"

type HouseBlend struct {
	description string
	cost        float64
}

func NewHouseBlend() *HouseBlend {
	return &HouseBlend{
		description: "House Blend Coffee",
		cost:        0.89,
	}
}

func (e *HouseBlend) GetDescription() string {
	return e.description
}

func (e *HouseBlend) Cost() float64 {
	return e.cost
}

var _ abstract.Beverage = (*HouseBlend)(nil)
