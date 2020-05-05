// Game initializes by showing the main menu.
// Afterwards, the player is presented by a menu of choices.
// Based on whose turn it is, a player may select an option.
// The following actions required by the option is then executed,
// and when it finishes, the action returns control to main(),
// causing the options-menu to reappear for the next player whose turn it is.
package main

import (
	"fmt"
	"os"

	"github.com/mradrianhh/go-fighter-game/commentary"
	"github.com/mradrianhh/go-fighter-game/game"
	"github.com/mradrianhh/go-fighter-game/models"
	"github.com/mradrianhh/go-fighter-game/persist"
	"github.com/mradrianhh/go-fighter-game/training"
)

var response int = 0
var charResponse = ""

const separator = "\n***\n"

func main() {
	fmt.Println(models.PlayerList[0].Name)
	models.PlayerList = models.AppendPlayer(models.PlayerList, models.NewPlayer("Adele", 3))
	fmt.Println(models.PlayerList[2].Name)
end:
	goto end
}

func start() {
	go models.Listen()
	showMainMenu()
}

// ShowMainMenu presents the main-menu to the user on the screen.
func showMainMenu() {
	for {
		fmt.Print(separator)
		fmt.Print("\nMain Menu\n")
		fmt.Print("\n1 - Start Game | 2 - Training | 3 - Save | 0 - Exit Game\n")
		if _, err := fmt.Scan(&response); err == nil {
			switch response {
			case 1:
				game.Start()
			case 2:
				training.Start()
			case 0:
				os.Exit(0)
			default:
				commentary.WrongInput()
			}
		}
	}
}

func save() error {
	if err := persist.Save("players", models.PlayerList); err != nil {
		return err
	}
	return nil
}

func load() error {
	if err := persist.Load("players", models.PlayerList); err != nil {
		return err
	}
	return nil
}
