package main

import (
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
}
