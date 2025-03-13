package weather

import (
	"fmt"
	"observer/observer"
)

type WeatherData struct {
	observer    map[int]observer.Observer
	temperature float32
	humidity    float32
	pressure    float32
}

func NewWeatherData() *WeatherData {
	return &WeatherData{
		observer: make(map[int]observer.Observer),
	}
}

func (w *WeatherData) RegisterObserver(observer observer.Observer) {
	w.observer[len(w.observer)] = observer
}

func (w *WeatherData) RemoveObserver(observer observer.Observer) {
	targetId := -1
	for k, v := range w.observer {
		if v == observer {
			targetId = k
			break
		}
	}
	if targetId != -1 {
		delete(w.observer, targetId)
	} else {
		fmt.Println("Observer not found")
	}
}

func (w *WeatherData) NotifyObservers() {
	for _, observer := range w.observer {
		observer.Update(w.temperature, w.humidity, w.pressure)
	}
}

func (w *WeatherData) MeasurementsChanged() {
	w.NotifyObservers()
}

func (w *WeatherData) SetMeasurements(temperature float32, humidity float32, pressure float32) {
	w.temperature = temperature
	w.humidity = humidity
	w.pressure = pressure
	w.MeasurementsChanged()
}

func (w *WeatherData) SayHi() {
	fmt.Println("Hello, I am WeatherData!")
}

// WeatherData implements observer.Subject interface
var _ observer.Subject = (*WeatherData)(nil)

func WeatherImplements() {
	// TODO
}
