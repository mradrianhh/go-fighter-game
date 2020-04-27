package game

import (
	"fmt"

	"github.com/mradrianhh/go-fighter-game/commentary"
)

// ShowFightMenu presents the user with the options he can take in combat.
func ShowFightMenu() {
	for {
		fmt.Print(separator)
		fmt.Printf("\n\tPlayer %d\n", currentPlayer.Number)
		fmt.Print(separator)
		fmt.Print("\nFight\n")
		fmt.Print("\n1 - Attack | 2 - Defend | 3 - Consume Auxiliary | 4 - Back\n")
		if _, err := fmt.Scan(&response); err == nil {
			switch response {
			case 1:
				currentPlayer.Attack(&opposingPlayer)
				goto end
			case 2:
				currentPlayer.Defend()
				goto end
			case 3:
				currentPlayer.ConsumeAuxiliaryItem()
				goto end
			case 4:
				goto end
			default:
				commentary.WrongInput()
			}
		}
	}
end:
	tick.Tick()
}
