package game

import "github.com/mradrianhh/go-fighter-game/models"

var response int = 0
var charResponse = ""

const separator = "\n***\n"

var shop models.Shop

// Player is the user playing the game.
var Player models.Player

func init() {
	shop = models.NewShop()

	Player = models.NewPlayer("Adrian")
}
