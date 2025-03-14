package weather

import (
	"observer/internal/util/observer"
)

type DisplayElement interface {
	Display()
}

type WeatherData struct {
	observer.Observable
	temperature float32
	humidity    float32
	pressure    float32
}

func NewWeatherData() *WeatherData {
	return &WeatherData{
		Observable: *observer.NewObservable(),
	}
}

func (w *WeatherData) MeasurementsChanged() {
	w.SetChanged()
	w.NotifyObservers(w.temperature, w.humidity, w.pressure)
}

func (w *WeatherData) SetMeasurements(temperature float32, humidity float32, pressure float32) {
	w.temperature = temperature
	w.humidity = humidity
	w.pressure = pressure
	w.MeasurementsChanged()
}

func (w *WeatherData) GetTemperature() float32 {
	return w.temperature
}

func (w *WeatherData) GetHumidity() float32 {
	return w.humidity
}

func (w *WeatherData) GetPressure() float32 {
	return w.pressure
}

// WeatherData implements observer.Subject interface
var _ observer.Subject = (*WeatherData)(nil)
