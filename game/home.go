package game

import (
	"fmt"

	"github.com/mradrianhh/go-fighter-game/commentary"
)

// Start initializes the necessary processes.
func Start() {
	showHome()
}

// ShowHome presents the user with the home menu.
func showHome() {
	for {
		currentPlayer.UpdateStats()
		fmt.Print(separator)
		fmt.Printf("\n\tPlayer %d\n", currentPlayer.Number)
		fmt.Print(separator)
		fmt.Print("\nHome\n")
		fmt.Print("\n1 - View Stats | 2 - View Shop | 3 - Fight | 0 - Exit to Main Menu\n")
		if _, err := fmt.Scan(&response); err == nil {
			switch response {
			case 1:
				ShowStats(&currentPlayer)
			case 2:
				ShowShop()
			case 3:
				ShowFightMenu()
			case 0:
				goto end
			default:
				commentary.WrongInput()
			}
		}
	}
end:
}
