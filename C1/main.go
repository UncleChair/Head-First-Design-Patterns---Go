package main

import (
	"duck/duck"
	"duck/rpg"
	"flag"
	"fmt"
)

func main() {
	mode := flag.String("mode", "duck", "运行模式: duck 或 rpg")
	flag.Parse()

	switch *mode {
	case "duck":
		duck.DuckImplements()
	case "rpg":
		rpg.RpgImplements()
	default:
		fmt.Println("Usage: run with -mode=duck or -mode=rpg")
	}
}
