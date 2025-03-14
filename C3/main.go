package main

import (
	"decorate/internal/abstract"
	"decorate/internal/drinks"
	"flag"
	"fmt"
)

func main() {
	mode := flag.String("mode", "starbazz", "运行模式: starbazz")
	flag.Parse()

	switch *mode {
	case "starbazz":
		StarbazzImplements()
	default:
		fmt.Println("Usage: run with -mode=starbazz")
	}
}

func StarbazzImplements() {
	fmt.Println("This is Decorator Pattern Implements in Go")
	var beverage abstract.Beverage

	// Espresso
	beverage = drinks.NewEspresso()
	fmt.Printf("%s $%.2f\n", beverage.GetDescription(), beverage.Cost())
	// Add double Mocha and Whip
	beverage = drinks.NewMocha(beverage)
	beverage = drinks.NewMocha(beverage)
	beverage = drinks.NewWhip(beverage)
	fmt.Printf("%s $%.2f\n", beverage.GetDescription(), beverage.Cost())
	// HouseBlend with Soy, Mocha, Whip
	beverage = drinks.NewHouseBlend()
	beverage = drinks.NewSoy(beverage)
	beverage = drinks.NewMocha(beverage)
	beverage = drinks.NewWhip(beverage)
	fmt.Printf("%s $%.2f\n", beverage.GetDescription(), beverage.Cost())
}
