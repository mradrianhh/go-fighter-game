package game

import (
	"github.com/mradrianhh/go-fighter-game/models"
)

var response int = 0
var charResponse = ""

const separator = "\n***\n"

var shop models.Shop

// Tick is the n-th loop cycle. It shows the total amount of cycles ran so far.
var Tick models.Tick

// CurrentPlayer is the player whose turn it currently is.
var currentPlayer models.Player

// Winner is the final player that won the game.
var Winner models.Player

// opposingPlayer is the player opposing the player whose turn it is.
var opposingPlayer models.Player

func init() {
	shop = models.NewShop()
	currentPlayer = models.PlayerList[0]
	opposingPlayer = models.PlayerList[1]
	Tick = models.NewTick([]func(){changeTurns})
	Winner = models.NewPlayer("Adrian", 1)
}
