package game

import (
	"fmt"
	"os"

	"github.com/mradrianhh/go-fighter-game/commentary"
)

// ShowMainMenu presents the main-menu to the user on the screen.
func ShowMainMenu() {
	for {
		fmt.Print(separator)
		fmt.Print("\nMain Menu\n")
		fmt.Print("\n1 - Start | 0 - Exit Game\n")
		if _, err := fmt.Scan(&response); err == nil {
			switch response {
			case 1:
				goto end // If the user wishes to start the game, we will jump out of the function and return control to main().
			case 0:
				os.Exit(0)
			default:
				commentary.WrongInput()
			}
		}
	}
end:
}
