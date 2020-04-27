package game

import (
	"fmt"
	"time"

	"github.com/mradrianhh/go-fighter-game/models"
)

var response int = 0
var charResponse = ""

const separator = "\n***\n"

var shop models.Shop

// Tick is the n-th loop cycle. It shows the total amount of cycles ran so far.
var tick models.Tick

// CurrentPlayer is the player whose turn it currently is.
var currentPlayer models.Player

// opposingPlayer is the player opposing the player whose turn it is.
var opposingPlayer models.Player

func init() {
	shop = models.NewShop()
	currentPlayer = models.PlayerList[0]
	opposingPlayer = models.PlayerList[1]
	tick = models.NewTick([]func(){changeTurns})
}

func changeTurns() {
	temp := currentPlayer
	currentPlayer = opposingPlayer
	opposingPlayer = temp
	fmt.Print(separator)
	fmt.Printf("\n\tPlayer %d's turn.\n", currentPlayer.Number)
	showPlayerInfo(currentPlayer)
	time.Sleep(1 * time.Second)
}
