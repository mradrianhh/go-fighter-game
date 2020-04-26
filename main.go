// Game initializes by showing the main menu.
// Afterwards, the player is presented by a menu of choices.
// Based on whose turn it is, a player may select an option.
// The following actions required by the option is then executed,
// and when it finishes, the action returns control to main(),
// causing the options-menu to reappear for the next player whose turn it is.
package main

import (
	"github.com/mradrianhh/go-fighter-game/game"
	"github.com/mradrianhh/go-fighter-game/models"
)

var shop models.Shop

// Player is the user playing the game.
var Player models.Player

func init() {
	shop = models.NewShop()
	Player = models.Player{Name: "Adrian", Damage: 10, Health: 100, Gold: 0}
}

func main() {
	game.ShowMainMenu()
	game.Loop()
}
