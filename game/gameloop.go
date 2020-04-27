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
		fmt.Print(separator)
		fmt.Print("\nOptions\n")
		fmt.Print("\n1 - View Stats | 2 - View Shop | 0 - Exit to Main Menu\n")
		if _, err := fmt.Scan(&response); err == nil {
			switch response {
			case 1:
				ShowStats(&Player)
			case 2:
				ShowShop()
			case 0:
				ShowMainMenu()
			default:
				commentary.WrongInput()
			}
		}
	}
}

// NextTurn increases "turn" by one and changes whether it's the players turn or not.
func NextTurn() {
	playersTurn = !playersTurn
	turn++
}
