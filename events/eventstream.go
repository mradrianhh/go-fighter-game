package events

// EventStream is a channel where one can push events into one end, and subscribers can receive them on the other.
var EventStream chan string = make(chan string)
