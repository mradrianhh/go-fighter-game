package game

import (
	"fmt"
	"time"

	"github.com/mradrianhh/go-fighter-game/events"
)

func changeTurns() {
	temp := currentPlayer
	currentPlayer = opposingPlayer
	opposingPlayer = temp
	go func() {
		events.EventStream <- "NextTurn"
	}()
	fmt.Print(separator)
	fmt.Printf("\n\tPlayer %d's turn.\n", currentPlayer.Number)
	showPlayerInfo(currentPlayer)
	time.Sleep(1 * time.Second)
}
