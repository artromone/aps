package main

import "time"

type Event struct {
	Type      string
	Data      interface{}
	Timestamp time.Time
}

func (eb *EventBus) Subscribe(eventType string, ch chan Event) {
	eb.subscribers[eventType] = append(eb.subscribers[eventType], ch)
}

type EventBus struct {
	subscribers map[string][]chan Event
}

func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[string][]chan Event),
	}
}

func (eb *EventBus) Publish(event Event) {
	if channels, exists := eb.subscribers[event.Type]; exists {
		for _, ch := range channels {
			ch <- event
		}
	}
}
