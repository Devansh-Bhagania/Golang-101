package main

import "fmt"

// Observer defines the interface for objects that should be notified of changes.
type Observer interface {
	Update(string)
}

// Subject maintains a list of observers and notifies them of any state changes.
type Subject struct {
	observers []Observer
}

func (s *Subject) Register(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *Subject) Notify(msg string) {
	for _, observer := range s.observers {
		observer.Update(msg)
	}
}

// ConcreteObserver is an implementation of the Observer interface.
type ConcreteObserver struct {
	ID string
}

func (o *ConcreteObserver) Update(msg string) {
	fmt.Printf("Observer %s received message: %s\n", o.ID, msg)
}

func main() {
	subject := &Subject{}

	observer1 := &ConcreteObserver{ID: "1"}
	observer2 := &ConcreteObserver{ID: "2"}

	subject.Register(observer1)
	subject.Register(observer2)

	subject.Notify("Event has occurred!")
}
