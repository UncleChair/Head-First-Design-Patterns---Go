package weather

type Subject interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers()
}

type Observer interface {
	Update(temp float32, humidity float32, pressure float32)
}

type DisplayElement interface {
	Display()
}

func WeatherImplements() {
	// TODO
}
