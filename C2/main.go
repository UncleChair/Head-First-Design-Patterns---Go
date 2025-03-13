package main

import (
	"flag"
	"fmt"
	"observer/weather"
)

func main() {
	mode := flag.String("mode", "weather", "运行模式: weather")
	flag.Parse()

	switch *mode {
	case "weather":
		weather.WeatherImplements()
	default:
		fmt.Println("Usage: run with -mode=weather")
	}
}
