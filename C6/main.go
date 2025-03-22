package main

import (
	"flag"
	"fmt"
)

func main() {
	mode := flag.String("mode", "command", "运行模式: command")
	flag.Parse()

	switch *mode {
	case "command":
		CommandImplements()
	default:
		fmt.Println("Usage: run with -mode=command")
	}
}

func CommandImplements() {
	fmt.Println("This is Command Pattern Implements in Go")
}
