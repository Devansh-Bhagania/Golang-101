package main

import "fmt"

// Notifier is the interface for sending notifications.
type Notifier interface {
	Send(message string)
}

// EmailNotifier is a concrete implementation of Notifier.
type EmailNotifier struct{}

func (e EmailNotifier) Send(message string) {
	fmt.Println("Sending Email:", message)
}

// SMSDecorator is a decorator that adds SMS functionality.
type SMSDecorator struct {
	Notifier
}

func (s SMSDecorator) Send(message string) {
	s.Notifier.Send(message)
	fmt.Println("Sending SMS:", message)
}

func main() {
	notifier := EmailNotifier{}
	notifierWithSMS := SMSDecorator{Notifier: notifier}

	notifierWithSMS.Send("Hello, World!")
}
