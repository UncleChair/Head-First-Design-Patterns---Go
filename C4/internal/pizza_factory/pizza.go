package pizza_factory

import "fmt"

type IPizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
	GetName() string
}

type Pizza struct {
	name     string
	dough    string
	sauce    string
	toppings []string
}

func (p *Pizza) Prepare() {
	fmt.Println("Preparing " + p.name)
	fmt.Println("Tossing dough...")
	fmt.Println("Adding sauce...")
	fmt.Println("Adding toppings:")
	for _, topping := range p.toppings {
		fmt.Println("  " + topping)
	}
}

func (p *Pizza) Bake() {
	fmt.Println("Bake for 25 minutes at 350")
}

func (p *Pizza) Cut() {
	fmt.Println("Cutting the pizza into diagonal slices")
}

func (p *Pizza) Box() {
	fmt.Println("Place pizza in official PizzaStore box")
}

func (p *Pizza) GetName() string {
	return p.name
}

type NYStyleCheesePizza struct {
	*Pizza
}

func CreateNYStyleCheesePizza() IPizza {
	return &NYStyleCheesePizza{
		Pizza: &Pizza{
			name:     "NY Style Sauce and Cheese Pizza",
			dough:    "Thin Crust Dough",
			sauce:    "Marinara Sauce",
			toppings: []string{"Grated Reggiano Cheese"},
		},
	}
}

type ChicagoStyleCheesePizza struct {
	*Pizza
}

func (p *ChicagoStyleCheesePizza) Cut() {
	fmt.Println("Cutting the pizza into square slices")
}

func NewChicagoStyleCheesePizza() IPizza {
	return &ChicagoStyleCheesePizza{
		Pizza: &Pizza{
			name:     "Chicago Style Deep Dish Cheese Pizza",
			dough:    "Extra Thick Crust Dough",
			sauce:    "Plum Tomato Sauce",
			toppings: []string{"Shredded Mozzarella Cheese"},
		},
	}
}
