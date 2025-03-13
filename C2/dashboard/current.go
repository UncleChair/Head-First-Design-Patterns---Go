package dashboard

import (
	"fmt"
	"observer/observer"
	"observer/weather"
)

type CurrentConditionDashboard struct {
	temperature float32
	humidity    float32
	weatherData *weather.WeatherData
}

func NewCurrentConditionDashboard(weatherData *weather.WeatherData) *CurrentConditionDashboard {
	return &CurrentConditionDashboard{weatherData: weatherData}
}

func (c *CurrentConditionDashboard) Display() {
	fmt.Printf("Current conditions: %f degrees and %f%% humidity\n", c.temperature, c.humidity)
}

func (c *CurrentConditionDashboard) Update(temperature float32, humidity float32, pressure float32) {
	c.temperature = temperature
	c.humidity = humidity
	c.Display()
}

// Make sure Dashboard implements required interfaces
var _ observer.Observer = (*CurrentConditionDashboard)(nil)
var _ observer.DisplayElement = (*CurrentConditionDashboard)(nil)
