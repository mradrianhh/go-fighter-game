package game

import (
	"fmt"

	"github.com/mradrianhh/go-fighter-game/commentary"
)

// Loop presents an options-menu of actions the user can take,
// and then passes control to selected action. Upon completion,
// loop() regains control, causing it to repeat.
func Loop() {
	for {
		fmt.Println("Options\t|\t1 - View Stats\t2 - View Shop\t")
		if _, err := fmt.Scan(&response); err == nil {
			switch response {
			case 1:
			case 2:
			default:
				commentary.WrongInput()
			}
		}
	}
}
