package game

import (
	"fmt"

	"github.com/mradrianhh/go-fighter-game/models"
)

var trainingOpponent models.Player = models.NewPlayer("Enemy")
var trainingPlayer models.Player = models.NewPlayer("Player")

// StartTraining orchestrates the initialization of the training mode.
func StartTraining() {

}

// ShowTrainingMenu presents the user with the actions to be taken in training mode.
func ShowTrainingMenu() {
	for {
		fmt.Print(separator)
		fmt.Print("\nTraining Menu\n")
		fmt.Print("\n1 - Attack | 2 - Defend | 3 - Consume Auxiliary | 0 - Exit to Main Menu\n")
		if _, err := fmt.Scan(&response); err == nil {
			switch response {
			case 1:
			case 2:
			case 3:
			case 0:
			}
		}
	}
}
