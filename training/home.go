package training

import (
	"fmt"

	"github.com/mradrianhh/go-fighter-game/commentary"
	"github.com/mradrianhh/go-fighter-game/models"
)

var trainingOpponent models.Player
var trainingPlayer models.Player

func init() {
	trainingOpponent = models.NewEnemyTrainer()
	trainingPlayer = models.NewPlayerTrainer()
}

// Start orchestrates the initialization of the training mode.
func Start() {
	configure()
	showHome()
}

// Configure sets all the necessary variables to their original state.
func configure() {
	trainingPlayer.Reset("ToTrainingPlayer")
	trainingOpponent.Reset("ToTrainingOpponent")
	playersTurn = true
	turn = 1
}

// ShowHome presents the user with the start menu in the training mode.
func showHome() {
	for {

		fmt.Print(separator)
		fmt.Print("\nTraining Menu\n")
		fmt.Print("\n1 - Scenario 1(Enemy attacks) | 2 - Scenario 2(Enemy defends) | 3 - Scenario 3(Enemy consumes auxiliary) | 0 - Exit to Main Menu\n")
		if _, err := fmt.Scan(&response); err == nil {
			switch response {
			case 1:
				runScenario1()
			case 2:
				runScenario2()
			case 3:
				runScenario3()
			case 0:
				goto end
			default:
				commentary.WrongInput()
			}
		}
	}
end:
}
