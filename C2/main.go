package main

import (
	"flag"
	"fmt"
	"observer/internal/pkg_instance/dashboard"
	"observer/internal/pkg_instance/weather"
)

func main() {
	mode := flag.String("mode", "weather", "运行模式: weather")
	flag.Parse()

	switch *mode {
	case "weather":
		WeatherImplements()
	default:
		fmt.Println("Usage: run with -mode=weather")
	}
}

func WeatherImplements() {
	weatherData := weather.NewWeatherData()
	CurrentConditionsDisplay := dashboard.NewCurrentConditionDashboard(weatherData)
	weatherData.RegisterObserver(CurrentConditionsDisplay)
	weatherData.SetMeasurements(20.5, 65, 1.0)
	weatherData.SetMeasurements(24.5, 55, 1.1)
	weatherData.SetMeasurements(22.5, 75, 0.9)
}
