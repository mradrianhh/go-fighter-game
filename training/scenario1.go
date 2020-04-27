package training

import (
	"fmt"
	"time"

	"github.com/mradrianhh/go-fighter-game/commentary"
	"github.com/mradrianhh/go-fighter-game/game"
)

// runScenario1 illustrates the opponent attacking.
func runScenario1() {
	for {
		fmt.Print(separator)
		fmt.Print("\nScenario 1\n")
		if playersTurn {
			fmt.Print("\n1 - Attack | 2 - Defend | 3 - Consume Auxiliary | 0 - Exit to Training Menu\n")
			if _, err := fmt.Scan(&response); err == nil {
				switch response {
				case 1:
					fmt.Print("\nAttacking...\n")
					trainingPlayer.Attack(&trainingOpponent)
					time.Sleep(1 * time.Second)
					nextTurn()
				case 2:
					fmt.Print("\nDefending...\n")
					trainingPlayer.Defend()
					time.Sleep(1 * time.Second)
					nextTurn()
				case 3:
					fmt.Print("\nConsuming...\n")
					trainingPlayer.ConsumeAuxiliaryItem()
					time.Sleep(1 * time.Second)
					nextTurn()
				case 0:
					goto end
				default:
					commentary.WrongInput()
				}
			}
			game.ShowStats(&trainingPlayer)
			game.ShowStats(&trainingOpponent)
		} else {
			fmt.Print("\nWaiting for opponent to make a move...\n")
			time.Sleep(1 * time.Second)
			trainingOpponent.Attack(&trainingPlayer)
			fmt.Print("\nAttacking...\n")
			time.Sleep(1 * time.Second)
			game.ShowStats(&trainingPlayer)
			game.ShowStats(&trainingOpponent)
			nextTurn()
		}
	}
end:
}
