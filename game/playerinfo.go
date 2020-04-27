package game

import (
	"fmt"

	"github.com/mradrianhh/go-fighter-game/models"
)

func showPlayerInfo(player models.Player) {
	fmt.Printf("\n\tGold: %d\n", player.Gold)
	fmt.Printf("\tDamage: %d\n", player.Damage)
	fmt.Printf("\tArmor: %d\n", player.Armor)
	fmt.Printf("\tHealth: %d\n", player.Health)
}
