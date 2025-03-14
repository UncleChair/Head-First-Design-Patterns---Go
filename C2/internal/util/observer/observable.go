package observer

import "fmt"

// Similar implementation to java.util.Observable
// Not that elegant, but it works :)
type Subject interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers(args ...any)
}

type Observer interface {
	Update(subject Subject, data ...interface{})
}

type Observable struct {
	changed   bool
	observers map[int]Observer
}

func (o *Observable) SetChanged() {
	o.changed = true
}

func (o *Observable) RegisterObserver(observer Observer) {
	o.observers[len(o.observers)] = observer
}

func (o *Observable) RemoveObserver(observer Observer) {
	targetId := -1
	for k, v := range o.observers {
		if v == observer {
			targetId = k
			break
		}
	}
	if targetId != -1 {
		delete(o.observers, targetId)
	} else {
		fmt.Println("Observer not found")
	}
}

func (o *Observable) NotifyObservers(args ...any) {
	if !o.changed {
		return
	}
	for _, observer := range o.observers {
		observer.Update(o, args...)
	}
	o.changed = false
}

var _ Subject = (*Observable)(nil)
