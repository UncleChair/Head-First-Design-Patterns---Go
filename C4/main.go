package main

import (
	"flag"
	"fmt"
)

func main() {
	mode := flag.String("mode", "pizza", "运行模式: pizza")
	flag.Parse()

	switch *mode {
	case "pizza":
		PizzaImplements()
	default:
		fmt.Println("Usage: run with -mode=pizza")
	}
}

func PizzaImplements() {
	fmt.Println("This is Decorator Pattern Implements in Go")
}
