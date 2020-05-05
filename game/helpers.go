package game

import (
	"fmt"
	"time"

	"github.com/mradrianhh/go-fighter-game/events"
	"github.com/mradrianhh/go-fighter-game/models"
)

func changeTurns() {
	for i := 0; i < len(models.PlayerList); i++ {
		if models.PlayerList[i].Stats.Health <= 0 {
			models.PlayerList = append(models.PlayerList[:i], models.PlayerList[i+1:]...)
		}
		length := len(models.PlayerList)
		if length == 1 {
			models.PlayerList[0] = Winner
			ShowEndscreen()
		}
	}
	temp := currentPlayer
	currentPlayer = opposingPlayer
	opposingPlayer = temp
	go func() {
		events.Notify(events.GlobalEventStream, "NextTurn")
	}()
	fmt.Print(separator)
	fmt.Printf("\n\tPlayer %d's turn.\n", currentPlayer.Number)
	showPlayerInfo(currentPlayer)
	time.Sleep(1 * time.Second)
}
