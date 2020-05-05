package events

import "fmt"

// GlobalEventListener ...
var GlobalEventListener EventListener = EventListener{}

func init() {
	GlobalEventListener.Events = make(map[string]func())
	GlobalEventListener.Events["Event 1"] = func() { fmt.Println("Event 1 pressed.") }
	GlobalEventListener.Events["Event 2"] = func() { fmt.Println("Event 2 pressed.") }
	GlobalEventListener.Events["Event 3"] = func() { fmt.Println("Event 3 pressed.") }
	go GlobalEventListener.Listen(GlobalEventStream)
}

// Listener ...
type Listener interface {
	Listen()
}

// EventListener listens for events.
type EventListener struct {
	Events map[string]func()
}

// Listen ...
func (eventListener *EventListener) Listen(eventStream <-chan string) {
	for {
		if event, ok := <-eventStream; ok {
			for e, f := range eventListener.Events {
				if event == e {
					f()
				}
			}
		}
	}
}
