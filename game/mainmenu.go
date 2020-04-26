package game

import (
	"fmt"
	"os"
)

// ShowMainMenu presents the main-menu to the user on the screen.
func ShowMainMenu() {
	for {
		fmt.Println("Main Menu | 1 - Start 2 - Exit")
		if _, err := fmt.Scan(&response); err == nil {
			switch response {
			case 1:
				goto end // If the user wishes to start the game, we will jump out of the function and return control to main().
			case 2:
				os.Exit(0)
			default:
				fmt.Println("Sorry, I can't understand. Try again.")
			}
		}
	}
end:
}
