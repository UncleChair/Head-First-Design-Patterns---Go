package dashboard

import (
	"fmt"
	"observer/internal/util/observer"
	"observer/internal/util_instance/weather"
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
	fmt.Printf("Current conditions: %.1f degrees and %.2f%% humidity\n", c.temperature, c.humidity)
}

func (c *CurrentConditionDashboard) Update(subject observer.Subject, data ...interface{}) {
	c.temperature = data[0].(float32)
	c.humidity = data[1].(float32)
	c.Display()
}

// Make sure Dashboard implements required interfaces
var _ observer.Observer = (*CurrentConditionDashboard)(nil)
var _ weather.DisplayElement = (*CurrentConditionDashboard)(nil)
