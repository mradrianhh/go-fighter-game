package events

// GlobalEventStream is a channel where one can push events into one end, and subscribers can receive them on the other.
var GlobalEventStream chan string = make(chan string)

// Notify notifies all the listeneres on the eventstream.
func Notify(eventStream chan<- string, event string) {
	eventStream <- event
}
