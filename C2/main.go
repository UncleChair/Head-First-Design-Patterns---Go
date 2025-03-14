package main

import (
	"flag"
	"fmt"
	"observer/internal/pkg_instance/dashboard"
	"observer/internal/pkg_instance/weather"
)

func main() {
	mode := flag.String("mode", "weather-util", "运行模式: weather-util, weather-implements")
	flag.Parse()

	switch *mode {
	case "weather-util":
		WeatherUtil()
	case "weather-implements":
		WeatherImplements()
	default:
		fmt.Println("Usage: run with -mode=weather-util, weather-implements")
	}
}

func WeatherImplements() {
	fmt.Println("This is Observer Pattern Self Implements in Go")
	weatherData := weather.NewWeatherData()
	CurrentConditionsDisplay := dashboard.NewCurrentConditionDashboard(weatherData)
	weatherData.RegisterObserver(CurrentConditionsDisplay)
	weatherData.SetMeasurements(20.5, 65, 1.0)
	weatherData.SetMeasurements(24.5, 55, 1.1)
	weatherData.SetMeasurements(22.5, 75, 0.9)
}

func WeatherUtil() {
	fmt.Println("This is Observer Pattern Mimic Java Util Observable in Go")
	weatherData := weather.NewWeatherData()
	CurrentConditionsDisplay := dashboard.NewCurrentConditionDashboard(weatherData)
	weatherData.RegisterObserver(CurrentConditionsDisplay)
	weatherData.SetMeasurements(20.5, 65, 1.0)
	weatherData.SetMeasurements(24.5, 55, 1.1)
	weatherData.SetMeasurements(22.5, 75, 0.9)
}
