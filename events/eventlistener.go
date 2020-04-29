package events

import (
	"fmt"
)

type dog struct {
	name    string
	sitters map[string][]chan string
}

func (d *dog) addSitter(event string, channel chan string) {
	if d.sitters == nil {
		d.sitters = make(map[string][]chan string)
	}

	if _, ok := d.sitters[event]; ok {
		d.sitters[event] = append(d.sitters[event], channel)
	} else {
		d.sitters[event] = []chan string{channel}
	}
}

func (d *dog) RemoveSitter(event string, channel chan string) {
	if _, ok := d.sitters[event]; ok {
		for i := range d.sitters[event] {
			if d.sitters[event][i] == channel {
				d.sitters[event] = append(d.sitters[event][:i], d.sitters[event][i+1:]...)
				break
			}
		}
	}
}

func (d *dog) emit(event string, response string) {
	if _, ok := d.sitters[event]; ok {
		for _, handler := range d.sitters[event] {
			go func(handler chan string) {
				handler <- response
			}(handler)
		}
	}
}

// Start ...
func Start() {
	balto := dog{"balto", nil}

	adrian := make(chan string)

	balto.addSitter("bark", adrian)

	go func() {
		for {
			msg := <-adrian
			fmt.Println("Adrian: " + msg)
		}
	}()

	fmt.Println("Balto barked!")
	balto.emit("bark", "Dont bark!")

	fmt.Scanln()
}
